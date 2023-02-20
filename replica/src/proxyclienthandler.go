package raxos

import (
	"fmt"
	"log"
	"os"
	"raxos/common"
	"raxos/proto/client"
	"strconv"
	"strings"
	"time"
)

// handler for new client batches

func (pr *Proxy) handleClientBatch(batch client.ClientBatch) {
	// put the client batch to the store
	pr.clientBatchStore.Add(batch)
	// add the batch id to the toBeProposed array
	pr.toBeProposed = append(pr.toBeProposed, batch.Id)

	if time.Now().Sub(pr.lastTimeProposed).Microseconds() >= pr.batchTime {
		if pr.lastProposedIndex-pr.committedIndex < pr.pipelineLength {
			proposeIndex := pr.lastProposedIndex + 1
			for proposeIndex+1 <= int64(len(pr.replicatedLog)) && pr.replicatedLog[proposeIndex].decided {
				proposeIndex++ // we always propose for a new index
			}
			msWait := int(pr.getLeaderWait(pr.getLeaderSequence(proposeIndex)))
			msWait = msWait * int(proposeIndex-pr.committedIndex) // adjust waiting for the pipelining
			msWait = msWait + pr.additionalDelay                  // for experiments
			if pr.instanceTimeouts[proposeIndex] != nil {
				pr.instanceTimeouts[proposeIndex].Cancel()
			}
			pr.debug("timeout for instance "+fmt.Sprintf("%v is %v", proposeIndex, msWait), 20)
			pr.instanceTimeouts[proposeIndex] = common.NewTimerWithCancel(time.Duration(msWait) * time.Microsecond)
			pr.instanceTimeouts[proposeIndex].SetTimeoutFuntion(func() {
				pr.proposeRequestIndex <- ProposeRequestIndex{index: proposeIndex}
			})
			pr.lastProposedIndex = proposeIndex
			pr.instanceTimeouts[proposeIndex].Start()
			if msWait != 0 && (pr.leaderMode == 1 || pr.leaderMode == 2) {
				pr.handleDecisionNotification()
			}
			pr.lastTimeProposed = time.Now()
		}
	}
}

// handler for client status request

func (pr *Proxy) handleClientStatus(status client.ClientStatus) {
	if status.Operation == 1 {
		if pr.serverStarted == false {
			// initiate gRPC connections
			pr.debug("proxy starting proposers  ", -1)
			pr.server.StartProposers()
			pr.serverStarted = true
			pr.startTime = time.Now()
		}
	}
	if status.Operation == 2 {
		pr.debug("proxy printing logs", 0)
		// print logs
		pr.printLog()
	}
	if status.Operation == 3 {
		pr.debug("proxy slowing down the proposing speed", 0)
		slowDown := status.Message
		split := strings.Split(slowDown, ",")
		for h := 0; h < len(split); h++ {
			splitItem := strings.Split(split[h], ":")
			nodeName, _ := strconv.Atoi(splitItem[0])
			if int64(nodeName) == pr.name {
				newDelay, _ := strconv.Atoi(splitItem[1])
				pr.additionalDelay = newDelay
				pr.debug("proxy slowing down the proposing speed by "+strconv.Itoa(pr.additionalDelay), 15)
				return
			}
		}
	}
}

// print the mempool and the consensus log to files

func (pr *Proxy) printLog() {
	pr.clientBatchStore.printStore(pr.logFilePath, pr.name) // print mem pool
	pr.printConsensusLog()                                  // print the replicated log
}

// print the replicated log to a file

func (pr *Proxy) printConsensusLog() {
	f, err := os.Create(pr.logFilePath + strconv.Itoa(int(pr.name)) + "-consensus.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < len(pr.replicatedLog); i++ {
		if pr.replicatedLog[i].committed == true {
			for j := 0; j < len(pr.replicatedLog[i].decidedBatch); j++ {
				batch, ok := pr.clientBatchStore.Get(pr.replicatedLog[i].decidedBatch[j])
				if !ok {
					panic("committed batch not in the store")
				} else {
					for k := 0; k < len(batch.Messages); k++ {
						_, _ = f.WriteString(strconv.Itoa(i) + "." + strconv.Itoa(j) + "." + strconv.Itoa(k) + ":" + batch.Messages[k].Message + "\n")
					}
				}
			}
		} else {
			break
		}
	}
}

// propose to index

func (pr *Proxy) proposeToIndex(proposeIndex int64) {

	if int64(len(pr.replicatedLog)) > proposeIndex && pr.replicatedLog[proposeIndex].decided == true {
		pr.debug("did not propose for index "+fmt.Sprintf("%v", proposeIndex)+" because it was decided", 9)
		return
	}
	pr.instanceTimeouts[proposeIndex] = nil

	pr.debug("proposing for index "+fmt.Sprintf("%v at time %v ms", proposeIndex, time.Now().Sub(pr.startTime).Milliseconds()), 20)

	if pr.leaderMode == 2 {
		if pr.isBeginningOfEpoch(proposeIndex) {
			pr.debug("proposing the last epoch summary for index "+fmt.Sprintf("%v", proposeIndex)+"", 13)
			pr.proposePreviousEpochSummary(proposeIndex)
			return
		}
	}

	batchSize := pr.batchSize
	if len(pr.toBeProposed) < batchSize {
		batchSize = len(pr.toBeProposed)
	}

	strProposals := make([]string, 0)
	btchProposals := make([]client.ClientBatch, 0)

	if batchSize == 0 {
		strProposals = []string{"nil"}
		btchProposals = append(btchProposals, client.ClientBatch{
			Sender:   -1,
			Messages: nil,
			Id:       "nil",
		})
		pr.debug("proposing empty values for index "+fmt.Sprintf("%v", proposeIndex), 9)
	} else {
		// send a new proposal Request to the ProposersChan
		strProposals = pr.toBeProposed[0:batchSize]
		pr.toBeProposed = pr.toBeProposed[batchSize:]
		btchProposals = make([]client.ClientBatch, 0)

		for i := 0; i < len(strProposals); i++ {
			btch, ok := pr.clientBatchStore.Get(strProposals[i])
			if !ok {
				panic("batch not found for the id")
			}
			btchProposals = append(btchProposals, btch)
		}
	}

	waitTime := int(pr.getLeaderWait(pr.getLeaderSequence(proposeIndex)))
	isLeader := true

	if pr.leaderMode == 3 {
		isLeader = false
	} else if pr.leaderMode != 3 && waitTime != 0 {
		isLeader = false
	}

	newProposalRequest := ProposeRequest{
		instance:             proposeIndex,
		proposalStr:          strProposals,
		proposalBtch:         btchProposals,
		isLeader:             isLeader,
		lastDecidedIndexes:   pr.lastDecidedIndexes,
		lastDecidedDecisions: pr.lastDecidedDecisions,
	}

	pr.proxyToProposerChan <- newProposalRequest
	pr.debug("proxy sent a proposal request to proposer  "+fmt.Sprintf("%v", newProposalRequest), -1)
	// create the slot index
	for len(pr.replicatedLog) < int(proposeIndex+1) {
		// create the new entry
		pr.replicatedLog = append(pr.replicatedLog, Slot{
			proposedBatch: nil,
			decidedBatch:  nil,
			decided:       false,
			committed:     false,
		})
	}

	pr.replicatedLog[proposeIndex] = Slot{
		proposedBatch: strProposals,
		decidedBatch:  pr.replicatedLog[proposeIndex].decidedBatch,
		decided:       pr.replicatedLog[proposeIndex].decided,
		committed:     pr.replicatedLog[proposeIndex].committed,
	}

	// reset the variables
	pr.lastDecidedIndexes = make([]int, 0)
	pr.lastDecidedDecisions = make([][]string, 0)
}
