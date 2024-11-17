package packets

import (
	"bytes"
	"encoding/binary"
)

type CreateAccessPointPacket struct {
	Header                    PacketHeader
	SecurityMode              uint16
	PassphraseSize            uint16
	Passphrase                [64]byte
	Username                  [33]byte
	Unknown1                  [15]byte
	LocalCommunicationID      uint64
	Reserved1                 uint16
	SceneID                   uint16
	Reserved2                 uint32
	Channel                   uint16
	NodeCountMax              uint8
	Reserved3                 uint8
	LocalCommunicationVersion uint16
	Reserved4                 [10]byte
	GameVersion               [16]byte
	PrivateIP                 [16]byte
	AddressFamily             int32
	ExternalProxyPort         uint16
	InternalProxyPort         uint16
}

func NewCreateAccessPointPacket(username string, gameVersion string, privateIP []byte) *CreateAccessPointPacket {
	p := &CreateAccessPointPacket{
		Header: PacketHeader{
			Magic:    MAGIC,
			PacketID: CreateAccessPoint,
			Version:  VERSION,
			DataSize: 192, // Total size of the packet data
		},
		LocalCommunicationID: 0x4200000000007E27,
		Channel:              3,
		NodeCountMax:         0xFF,
		AddressFamily:        2,
		ExternalProxyPort:    30456,
		InternalProxyPort:    31456,
	}

	copy(p.Username[:], []byte("RyuDoctor"))
	copy(p.GameVersion[:], []byte("0.1.0-TEST"))
	copy(p.PrivateIP[:], []byte{127, 0, 0, 1})

	return p
}

func (p *CreateAccessPointPacket) GetHeader() *PacketHeader {
	return &p.Header
}

func (p *CreateAccessPointPacket) Encode() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, ByteOrder, p)
	return buf.Bytes()
}

func (p *CreateAccessPointPacket) Decode(data []byte) error {
	return binary.Read(bytes.NewReader(data), ByteOrder, p)
}
