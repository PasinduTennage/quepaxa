package raxos

import (
	"math/rand"
	"raxos/proto"
	"strconv"
)

/*
	Recorder: handler for the recorder. All the recorder logic goes here
*/

func (in *Instance) handleRecorderConsensusMessage(consensusMessage *proto.GenericConsensus) {
	/*
		We use pipelining, so it is possible that the recorder received messages corresponding to different instances without order
	*/
	in.recorderReplicatedLog = in.initializeSlot(in.recorderReplicatedLog, consensusMessage.Index) // create the slot if not already created

	// case 1: if this entry has been previously decided, send the response back
	if in.recorderReplicatedLog[consensusMessage.Index].decided == true && consensusMessage.M != in.decideMessage {
		in.debug("Received a message from "+strconv.Itoa(int(consensusMessage.Sender))+" that is not a decide message, replying with a decision message", 1)
		in.sendRecorderDecided(consensusMessage)
		return
	}

	// case 2: if this is decide message, and if I have not previously decided on this slot, decide it. This case is already handled in the common.go

	// case 3: other consensus messages from a higher step
	if in.recorderReplicatedLog[consensusMessage.Index].S < consensusMessage.S {
		in.recorderReplicatedLog[consensusMessage.Index].S = consensusMessage.S

		// case 3.1 a propose message
		if consensusMessage.M == in.proposeMessage {
			in.handleRecorderProposeMessage(consensusMessage)
		}
	}

	// case 4 a spreadE message

	if consensusMessage.M == in.spreadEMessage && in.recorderReplicatedLog[consensusMessage.Index].S == consensusMessage.S {
		in.handleRecorderSpreadEMessage(consensusMessage)
	}

	// case 5 a spreadCgatherE message

	if consensusMessage.M == in.spreadCgatherEMessage && in.recorderReplicatedLog[consensusMessage.Index].S == consensusMessage.S {
		in.handleRecorderSpreadCGatherEMessage(consensusMessage)
	}

	// send the response back to proposer

	consensusReply := proto.GenericConsensus{
		Sender:   in.nodeName,
		Receiver: consensusMessage.Sender,
		Index:    consensusMessage.Index,
		M:        consensusMessage.M,
		S:        in.recorderReplicatedLog[consensusMessage.Index].S,
		P: &proto.GenericConsensusValue{
			Id:  in.recorderReplicatedLog[consensusMessage.Index].P.id,
			Fit: in.recorderReplicatedLog[consensusMessage.Index].P.fit,
		},
		E: in.getGenericConsensusValueArray(in.recorderReplicatedLog[consensusMessage.Index].E),
		C: in.getGenericConsensusValueArray(in.recorderReplicatedLog[consensusMessage.Index].C),
		D: in.recorderReplicatedLog[consensusMessage.Index].decided,
		DS: &proto.GenericConsensusValue{
			Id:  in.recorderReplicatedLog[consensusMessage.Index].decision.id,
			Fit: in.recorderReplicatedLog[consensusMessage.Index].decision.fit,
		},
		PR:          in.recorderReplicatedLog[consensusMessage.Index].proposer,
		Destination: in.consensusMessageProposerDestination,
	}

	rpcPair := RPCPair{
		Code: in.genericConsensusRpc,
		Obj:  &consensusReply,
	}
	in.debug("Recorder sending a generic consensus reply message to "+strconv.Itoa(int(consensusMessage.Sender)), 1)

	in.sendMessage(consensusMessage.Sender, rpcPair)

}

/*
	handler for spreadCGatherE message
*/

func (in *Instance) handleRecorderSpreadCGatherEMessage(consensusMessage *proto.GenericConsensus) {
	for i := 0; i < len(consensusMessage.C); i++ {
		in.recorderReplicatedLog[consensusMessage.Index].C = append(in.recorderReplicatedLog[consensusMessage.Index].C, Value{
			id:  consensusMessage.C[i].Id,
			fit: consensusMessage.C[i].Fit,
		})
	}
	in.debug("Recorder processed a SpreadCGatherE message and updated the C set", 1)
}

/*
	handler for spreadE message
*/

func (in *Instance) handleRecorderSpreadEMessage(consensusMessage *proto.GenericConsensus) {
	for i := 0; i < len(consensusMessage.E); i++ {
		in.recorderReplicatedLog[consensusMessage.Index].E = append(in.recorderReplicatedLog[consensusMessage.Index].E, Value{
			id:  consensusMessage.E[i].Id,
			fit: consensusMessage.E[i].Fit,
		})
	}
	in.debug("Recorder processed a spreadE message and added the E set to its E set", 1)
}

/*
	handler for propose messages
*/

func (in *Instance) handleRecorderProposeMessage(consensusMessage *proto.GenericConsensus) {
	if consensusMessage.Sender == in.getDeterministicLeader1() {
		in.recorderReplicatedLog[consensusMessage.Index].P = Value{
			id:  consensusMessage.P.Id,
			fit: strconv.FormatInt(in.Hi, 10) + "." + strconv.FormatInt(consensusMessage.Sender, 10),
		}
	} else {
		in.recorderReplicatedLog[consensusMessage.Index].P = Value{
			id:  consensusMessage.P.Id,
			fit: string(rand.Intn(int(in.Hi-10))) + "." + strconv.FormatInt(consensusMessage.Sender, 10),
		}
	}
	in.recorderReplicatedLog[consensusMessage.Index].E = append(in.recorderReplicatedLog[consensusMessage.Index].E, in.recorderReplicatedLog[consensusMessage.Index].P)
	in.debug("Recorder assigned the priority to the proposal and appended to E set", 1)
}

/*
	mark the slot as decided
*/

func (in *Instance) recordRecorderDecide(consensusMessage *proto.GenericConsensus) {
	in.recorderReplicatedLog[consensusMessage.Index].decided = true
	in.recorderReplicatedLog[consensusMessage.Index].S = consensusMessage.S
	in.recorderReplicatedLog[consensusMessage.Index].decision = Value{
		id:  consensusMessage.DS.Id,
		fit: consensusMessage.DS.Fit,
	}
	in.recorderReplicatedLog[consensusMessage.Index].proposer = consensusMessage.PR

	in.recorderReplicatedLog[consensusMessage.Index].P = Value{}
	in.recorderReplicatedLog[consensusMessage.Index].E = []Value{}
	in.recorderReplicatedLog[consensusMessage.Index].C = []Value{}
	in.recorderReplicatedLog[consensusMessage.Index].U = []Value{}
	in.recorderReplicatedLog[consensusMessage.Index].proposeResponses = []*proto.GenericConsensus{}
	in.recorderReplicatedLog[consensusMessage.Index].spreadEResponses = []*proto.GenericConsensus{}
	in.recorderReplicatedLog[consensusMessage.Index].spreadCGatherEResponses = []*proto.GenericConsensus{}
	in.recorderReplicatedLog[consensusMessage.Index].gatherCResponses = []*proto.GenericConsensus{}

}

/*
	send a decide message to the proposer
*/

func (in *Instance) sendRecorderDecided(consensusMessage *proto.GenericConsensus) {
	consensusReply := proto.GenericConsensus{
		Sender:   in.nodeName,
		Receiver: consensusMessage.Sender,
		Index:    consensusMessage.Index,
		M:        in.decideMessage,
		S:        in.recorderReplicatedLog[consensusMessage.Index].S,
		P:        nil,
		E:        nil,
		C:        nil,
		D:        true,
		DS: &proto.GenericConsensusValue{
			Id:  in.recorderReplicatedLog[consensusMessage.Index].decision.id,
			Fit: in.recorderReplicatedLog[consensusMessage.Index].decision.fit,
		},
		PR:          in.recorderReplicatedLog[consensusMessage.Index].proposer,
		Destination: in.consensusMessageCommonDestination,
	}

	rpcPair := RPCPair{
		Code: in.genericConsensusRpc,
		Obj:  &consensusReply,
	}
	in.debug("sending a decide consensus message to "+strconv.Itoa(int(consensusMessage.Sender)), 1)

	in.sendMessage(consensusMessage.Sender, rpcPair)
}
