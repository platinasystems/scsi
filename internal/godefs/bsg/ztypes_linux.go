// Code generated by scsi_id/godefs; DO NOT EDIT.

package bsg

const (
	BSG_PROTOCOL_SCSI = 0x0

	BSG_SUB_PROTOCOL_SCSI_CMD       = 0x0
	BSG_SUB_PROTOCOL_SCSI_TMF       = 0x1
	BSG_SUB_PROTOCOL_SCSI_TRANSPORT = 0x2

	BSG_FLAG_Q_AT_TAIL = 0x10
	BSG_FLAG_Q_AT_HEAD = 0x20
)

type SgIoV4 struct {
	Guard           int32
	Protocol        uint32
	Subprotocol     uint32
	RequestLen      uint32
	Request         uint64
	RequestTag      uint64
	RequestAttr     uint32
	RequestPriority uint32
	RequestExtra    uint32
	MaxResponseLen  uint32
	Response        uint64
	DoutIovecCount  uint32
	DoutXferLen     uint32
	DinIovecCount   uint32
	DinXferLen      uint32
	DoutXferp       uint64
	DinXferp        uint64
	Timeout         uint32
	Flags           uint32
	UsrPtr          uint64
	SpareIn         uint32
	DriverStatus    uint32
	TransportStatus uint32
	DeviceStatus    uint32
	RetryDelay      uint32
	Info            uint32
	Duration        uint32
	ResponseLen     uint32
	DinResid        int32
	DoutResid       int32
	GeneratedTag    uint64
	SpareOut        uint32
	Padding         uint32
}