package messages

import (
	"errors"

	kyberk2so "github.com/symbolicsoft/kyber-k2so"
)

var (
	MagicNumber     = [4]byte{'K', 'B', 'P', 'M'}
	SizeMagicNumber = len(MagicNumber)
)

type Type uint16

const (
	TypeUnknown Type = iota
	TypePubkey
	TypeGeneral
)

type MsgPubkey struct {
	Pubkey *[kyberk2so.Kyber512PKBytes]byte
}

func (m *MsgPubkey) Marshal() ([]byte, error) {
	buf := make([]byte, kyberk2so.Kyber512PKBytes+len(MagicNumber))

	copy(buf[0:len(MagicNumber)], MagicNumber[:])
	copy(buf[len(MagicNumber):], m.Pubkey[:])

	return buf, nil
}

type MsgGeneral struct {
	Ciphertext *[kyberk2so.Kyber512CTBytes]byte
	Nonce      *[]byte
	Data       *[]byte
}

func (m *MsgGeneral) Marshal() ([]byte, error) {
	if m.Nonce == nil {
		return nil, errors.New("nil nonce")
	}

	if m.Data == nil {
		return nil, errors.New("no data to marshal")
	}

	buf := make([]byte, SizeMagicNumber+kyberk2so.Kyber512CTBytes+len(*m.Nonce)+len(*m.Data))

	copy(buf[0:SizeMagicNumber], MagicNumber[:])
	copy(buf[SizeMagicNumber:], m.Ciphertext[:])
	copy(buf[SizeMagicNumber+kyberk2so.Kyber512CTBytes:], *m.Nonce)
	copy(buf[SizeMagicNumber+kyberk2so.Kyber512CTBytes+len(*m.Nonce):], *m.Data)

	return buf, nil
}
