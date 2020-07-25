package packets

import (
	"bytes"
	"fmt"
	"io"
)

// Pubrec represents the MQTT Pubrec  packet.
type Pubrec struct {
	Version   Version
	FixHeader *FixHeader
	PacketID  PacketID

	// V5
	Code       byte
	Properties *Properties
}

func (p *Pubrec) String() string {
	return fmt.Sprintf("Pubrec, Pid: %v", p.PacketID)
}

// NewPubrecPacket returns a Pubrec instance by the given FixHeader and io.Reader.
func NewPubrecPacket(fh *FixHeader, version Version, r io.Reader) (*Pubrec, error) {
	p := &Pubrec{FixHeader: fh, Version: version}
	if p.Version == Version5 {
		p.Properties = &Properties{}
	}
	err := p.Unpack(r)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// NewPubrel returns the Pubrel struct related to the Pubrec struct in QoS 2.
func (p *Pubrec) NewPubrel() *Pubrel {
	pub := &Pubrel{FixHeader: &FixHeader{PacketType: PUBREL, Flags: FlagPubrel}, Version: p.Version}
	pub.PacketID = p.PacketID
	return pub
}

// Pack encodes the packet struct into bytes and writes it into io.Writer.
func (p *Pubrec) Pack(w io.Writer) error {
	p.FixHeader = &FixHeader{PacketType: PUBREC, Flags: RESERVED}
	bufw := &bytes.Buffer{}
	writeUint16(bufw, p.PacketID)
	bufw.WriteByte(p.Code)
	if p.Properties != nil {
		p.Properties.Pack(bufw, PUBREC)
	}
	p.FixHeader.RemainLength = bufw.Len()
	err := p.FixHeader.Pack(w)
	if err != nil {
		return err
	}
	_, err = bufw.WriteTo(w)
	return err
}

// Unpack read the packet bytes from io.Reader and decodes it into the packet struct.
func (p *Pubrec) Unpack(r io.Reader) error {
	restBuffer := make([]byte, p.FixHeader.RemainLength)
	_, err := io.ReadFull(r, restBuffer)
	if err != nil {
		return errMalformed(err)
	}
	bufr := bytes.NewBuffer(restBuffer)
	p.PacketID, err = readUint16(bufr)
	if err != nil {
		return err
	}
	if p.Version == Version5 {
		p.Properties = &Properties{}
		if p.Code, err = bufr.ReadByte(); err != nil {
			return err
		}
		if !ValidateCode(PUBREC, p.Code) {
			return protocolErr(invalidReasonCode(p.Code))
		}
		if err := p.Properties.Unpack(bufr, PUBREC); err != nil {
			return err
		}
	}
	return nil
}
