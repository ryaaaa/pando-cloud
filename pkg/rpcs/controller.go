package rpcs

import (
	"github.com/PandoCloud/pando-cloud/pkg/protocol"
	"github.com/PandoCloud/pando-cloud/pkg/tlv"
)

type ArgsOnStatus struct {
	DeviceId  uint64
	Timestamp uint64
	Subdata   []protocol.SubData
}
type ReplyOnStatus ReplyEmptyResult

type ArgsOnEvent struct {
	DeviceId  uint64
	TimeStamp uint64
	SubDevice uint16
	No        uint16
	Priority  uint16
	Params    []tlv.TLV
}
type ReplyOnEvent ReplyEmptyResult
