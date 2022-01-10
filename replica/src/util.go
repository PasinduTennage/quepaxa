package raxos

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"raxos/proto"
)

func (in *Instance) debug(message string) {
	fmt.Printf("%s\n", message)

}

func getRealSizeOf(v interface{}) (int, error) {
	b := new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(v); err != nil {
		return 0, err
	}
	return b.Len(), nil
}

func getStringOfSizeN(length int) string {
	str := "a"
	size, _ := getRealSizeOf(str)
	for size < length {
		str = str + "a"
		size, _ = getRealSizeOf(str)
	}
	return str
}

func (in *Instance) getNewCopyOfMessage(code uint8, msg proto.Serializable) proto.Serializable {

	if code == in.clientRequestBatchRpc {

		clientRequestBatch := msg.(*proto.ClientRequestBatch)
		return &proto.ClientRequestBatch{
			Sender:   clientRequestBatch.Sender,
			Requests: clientRequestBatch.Requests,
			Id:       clientRequestBatch.Id,
		}

	}

	if code == in.clientResponseBatchRpc {
		clientResponseBatch := msg.(*proto.ClientResponseBatch)
		return &proto.ClientResponseBatch{
			Receiver:  clientResponseBatch.Receiver,
			Responses: clientResponseBatch.Responses,
			Id:        clientResponseBatch.Id,
		}

	}

	if code == in.genericConsensusRpc {

		genericConsensus := msg.(*proto.GenericConsensus)
		return &proto.GenericConsensus{
			Index:       genericConsensus.Index,
			M:           genericConsensus.M,
			S:           genericConsensus.S,
			P:           genericConsensus.P,
			E:           genericConsensus.E,
			C:           genericConsensus.C,
			D:           genericConsensus.D,
			DS:          genericConsensus.DS,
			PR:          genericConsensus.PR,
			Destination: genericConsensus.Destination,
		}

	}

	if code == in.MessageBlockRpc {

		messageBlock := msg.(*proto.MessageBlock)
		return &proto.MessageBlock{
			Hash:     messageBlock.Hash,
			Requests: messageBlock.Requests,
		}

	}

	if code == in.MessageBlockRequestRpc {
		messageBlockRequest := msg.(*proto.MessageBlockRequest)
		return &proto.MessageBlockRequest{Hash: messageBlockRequest.Hash}
	}

	return nil

}
