package dnsboot

import (
	"encoding/binary"
	"errors"
	"time"

	"github.com/u6du/go-rfc1924/base85"
	key "github.com/u6du/key/ed25519"
	"golang.org/x/crypto/ed25519"
)

var (
	ErrVerify  = errors.New("verify")
	ErrTimeout = errors.New("timeout")
	ErrDecode  = errors.New("decode")
	ErrEmpty   = errors.New("empty")
)

const TimeOutHour = uint32(1)

func Decode(txt string) ([]byte, error) {
	b, err := base85.DecodeString(txt)
	return b, err
}

func Verify(txt string) ([]byte, error) {
	if len(txt) > ed25519.SignatureSize {
		b, err := Decode(txt)
		if err != nil {
			return []byte{}, ErrDecode
		}
		n := ed25519.SignatureSize
		ctx := b[n:]
		sign := b[:n]
		if ed25519.Verify(key.GodPublic, ctx, sign) {

			hour := ctx[0:4]
			ctx := ctx[4:]

			cost := uint32(time.Now().Unix()/3600) - binary.LittleEndian.Uint32(hour)
			var state error
			if cost >= TimeOutHour {
				state = ErrTimeout
			} else {
				state = nil
			}
			return ctx, state

		} else {
			return []byte{}, ErrVerify
		}
	}
	return []byte{}, ErrEmpty
}
