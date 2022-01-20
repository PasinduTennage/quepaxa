package cmd

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"raxos/proto"
)

/*
	iteratively calculates the number of elements in the 2d array
*/
func (cl *Client) getNumberofRequests(requests [][]sentRequestBatch) int {
	count := 0
	for i := 0; i < numRequestGenerationThreads; i++ {
		count += len(requests[i])
	}
	return count
}

/*
	returns the matching response batch corresponding to the request batch id
*/

func (cl *Client) getMatchingResponseBatch(id string) int {
	for i := 0; i < len(cl.receivedResponses); i++ {
		if cl.receivedResponses[i].batch.Id == id {
			return i
		}
	}
	return -1
}

/*
	Add value N to list, M times
*/

func (cl *Client) addValueNToArrayMTimes(list []int64, N int64, M int) []int64 {
	for i := 0; i < M; i++ {
		list = append(list, N)
	}
	return list
}

/*
	Maps the request with the response batch
    Compute the time taken for each request
	Computer the error rate
	Compute the throughput as successfully committed requests per second (doesn't include failed requests)
	Compute the latency by including the timeout requests
	Prints the basic stats to the stdout
*/

func (cl *Client) computeStats() {

	numTotalRequests := cl.getNumberofRequests(cl.sentRequests)
	var latencyList []int64    // contains the time duration spent for each request in micro seconds (includes failed requests)
	var throughputList []int64 // contains the time duration spent for successful requests
	for i := 0; i < numRequestGenerationThreads; i++ {
		for j := 0; j < len(cl.sentRequests[i]); j++ {
			batch := cl.sentRequests[i][j]
			batchId := batch.batch.Id
			matchingResponseIndex := cl.getMatchingResponseBatch(batchId)
			if matchingResponseIndex == -1 {
				// there is no response for this batch of requests, hence they are considered as failed
				cl.addValueNToArrayMTimes(latencyList, cl.replicaTimeout*1000, len(batch.batch.Requests))
				cl.printRequests(batch.batch, batch.time.Sub(cl.startTime).Microseconds(), batch.time.Sub(cl.startTime).Microseconds()+cl.replicaTimeout*1000)
			} else {
				responseBatch := cl.receivedResponses[matchingResponseIndex]
				startTime := batch.time
				endTime := responseBatch.time
				batchLatency := endTime.Sub(startTime).Microseconds()
				cl.addValueNToArrayMTimes(latencyList, batchLatency, len(batch.batch.Requests))
				cl.addValueNToArrayMTimes(throughputList, batchLatency, len(batch.batch.Requests))
				cl.printRequests(batch.batch, batch.time.Sub(cl.startTime).Microseconds(), endTime.Sub(cl.startTime).Microseconds())
			}

		}
	}

	medianLatency, _ := stats.Median(cl.getFloat64List(latencyList))
	percentile99, _ := stats.Percentile(cl.getFloat64List(latencyList), 99.0) // tail latency
	duration := cl.testDuration
	errorRate := (numTotalRequests - len(throughputList)) * 100.0 / numTotalRequests

	fmt.Printf("Total Requests (includes timeout requests) := %v \n", numTotalRequests)
	fmt.Printf("Total time := %v seconds\n", duration)
	fmt.Printf("Throughput (successfully committed requests) := %v requests per second\n", len(throughputList)/int(duration))
	fmt.Printf("Median Latency (includes timeout requests) := %v micro seconds per request\n", medianLatency)
	fmt.Printf("99 pecentile latency (includes timeout requests) := %v micro seconds per request\n", percentile99)
	fmt.Printf("Error Rate := %v \n", float64(errorRate))
}

/*
	Converts int[] to float64[]
*/

func (cl *Client) getFloat64List(list []int64) []float64 {
	var array []float64
	for i := 0; i < len(list); i++ {
		array = append(array, float64(list[i]))
	}
	return array
}

/*
	Print a client request batch with arrival time and end time
*/

func (cl *Client) printRequests(messages proto.ClientRequestBatch, startTime int64, endTime int64) {
	fmt.Print(messages.Id, "\n")
	for i := 0; i < len(messages.Requests); i++ {
		fmt.Print(messages.Requests[i].Message, ",", startTime, ",", endTime, "\n")
	}
}
