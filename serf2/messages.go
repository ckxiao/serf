package serf2

import (
	"bytes"
	"github.com/ugorji/go/codec"
)

// messageType are the types of gossip messages Serf will send along
// memberlist.
type messageType uint8

const (
	messageLeaveType messageType = iota
	messageRemoveFailedType
)

// messageLeave is the message broadcasted to signal the intentional to
// leave.
type messageLeave struct {
	Node string
}

// messageRemoveFailed is the message broadcasted to force remove
// a failed node from the failed node list and the member list.
type messageRemoveFailed struct {
	Node string
}

func decodeMessage(buf []byte, out interface{}) error {
	var handle codec.MsgpackHandle
	return codec.NewDecoder(bytes.NewBuffer(buf), &handle).Decode(out)
}

func encodeMessage(t messageType, msg interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(uint8(t))

	handle := codec.MsgpackHandle{}
	encoder := codec.NewEncoder(buf, &handle)
	err := encoder.Encode(msg)
	return buf.Bytes(), err
}