package packets

import (
	"bytes"
	"encoding/binary"
)

type InitializePacket struct {
	Header     PacketHeader
	ClientID   [16]byte
	MacAddress [6]byte
}

func NewInitializePacket() *InitializePacket {
	return &InitializePacket{
		Header: PacketHeader{
			Magic:    MAGIC,
			PacketID: Initialize,
			Version:  VERSION,
			DataSize: 22, // 16 + 6 bytes
		},
	}
}

func (p *InitializePacket) GetHeader() *PacketHeader {
	return &p.Header
}

func (p *InitializePacket) Encode() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, ByteOrder, &p.Header)
	binary.Write(buf, ByteOrder, &p.ClientID)
	binary.Write(buf, ByteOrder, &p.MacAddress)
	return buf.Bytes()
}

func (p *InitializePacket) Decode(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, ByteOrder, &p.Header); err != nil {
		return err
	}
	if err := binary.Read(buf, ByteOrder, &p.ClientID); err != nil {
		return err
	}
	if err := binary.Read(buf, ByteOrder, &p.MacAddress); err != nil {
		return err
	}
	return nil
}
