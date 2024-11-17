package ldnhealthcheck

import (
	"RyuLdnWebsite/ldnhealthcheck/packets"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type RyujinxLdnClient struct {
	conn net.Conn
}

func NewRyujinxLdnClient(host string, port int, timeout time.Duration) (*RyujinxLdnClient, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		return nil, err
	}
	return &RyujinxLdnClient{conn: conn}, nil
}

func (c *RyujinxLdnClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *RyujinxLdnClient) Send(packet packets.Packet) error {
	data := packet.Encode()
	_, err := c.conn.Write(data)
	return err
}

func (c *RyujinxLdnClient) Receive() (packets.Packet, error) {
	headerBuf := make([]byte, binary.Size(packets.PacketHeader{}))
	_, err := c.conn.Read(headerBuf)
	if err != nil {
		return nil, err
	}

	var header packets.PacketHeader
	if err := binary.Read(bytes.NewReader(headerBuf), packets.ByteOrder, &header); err != nil {
		return nil, err
	}

	dataBuf := make([]byte, header.DataSize)
	_, err = c.conn.Read(dataBuf)
	if err != nil {
		return nil, err
	}

	var packet packets.Packet
	switch header.PacketID {
	case packets.Initialize:
		p := packets.NewInitializePacket()
		packet = p
	default:
		return nil, fmt.Errorf("unknown packet type: %d", header.PacketID)
	}

	fullData := append(headerBuf, dataBuf...)
	if err := packet.Decode(fullData); err != nil {
		return nil, err
	}

	return packet, nil
}
