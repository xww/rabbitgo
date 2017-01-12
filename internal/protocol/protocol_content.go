package protocol

type MsgProtocol struct {
	Type uint8
	Version uint16
	Token uint32
	MsgLength uint16
	MsgContent string
	CheckSum uint32
}

type CmdProtocol struct {
	Type uint8
	Version uint16
	Token uint32
	MsgLength uint16
	MsgContent string
	CheckSum uint32
}

type HeartbeatProtocol struct {
	Type uint8
	Version uint16
	Token uint32
	MsgLength uint16
	MsgContent string
	CheckSum uint32
}
