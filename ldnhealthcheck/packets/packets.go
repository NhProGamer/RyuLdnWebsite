package packets

import "encoding/binary"

const (
	MAGIC    = 1313098834
	VERSION  = 1
	MAX_SIZE = 131072
)

var ByteOrder = binary.LittleEndian

type PacketID uint8

const (
	Initialize PacketID = iota
	Passphrase
	CreateAccessPoint
	CreateAccessPointPrivate
	ExternalProxy
	ExternalProxyToken
	ExternalProxyState
	SyncNetwork
	Reject
	RejectReply
	Scan
	ScanReply
	ScanReplyEnd
	Connect
	ConnectPrivate
	Connected
	Disconnect
	ProxyConfig
	ProxyConnect
	ProxyConnectReply
	ProxyData
	ProxyDisconnect
	SetAcceptPolicy
	SetAdvertiseData
	Ping         PacketID = 254
	NetworkError PacketID = 255
)

type PacketHeader struct {
	Magic    uint32
	PacketID PacketID
	Version  uint8
	_        uint16 // padding
	DataSize uint32
}

type Packet interface {
	Encode() []byte
	Decode([]byte) error
	GetHeader() *PacketHeader
}
