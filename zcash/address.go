// zcash provides a few tools to manipulate Zcash z-addresses.
package zcash

import (
	"crypto/rand"
	"errors"

	"github.com/FiloSottile/zcash-mini/sha256"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/curve25519"
)

var (
	ProdSpendingKey = [2]byte{0xAB, 0x36}
	TestSpendingKey = [2]byte{0xAC, 0x08}
	ProdAddress     = [2]byte{0x16, 0x9A}
	TestAddress     = [2]byte{0x16, 0xB6}
	ViewingKey      = [2]byte{0, 0} //Is this specified yet?
)

var (
	ErrChecksum      = errors.New("checksum error")
	ErrInvalidFormat = errors.New("invalid format: version and/or checksum bytes missing")
	ErrInvalidKey    = errors.New("invalid key: must be 32 bytes with first 4 bits set to 0")
)

// Base58Decode decodes a Base58Check encoding with two version bytes.
func Base58Decode(s string) (result []byte, version [2]byte, err error) {
	decoded, vByte, err := base58.CheckDecode(s)
	switch err {
	case base58.ErrChecksum:
		err = ErrChecksum
		return
	case base58.ErrInvalidFormat:
		err = ErrInvalidFormat
		return
	default:
		return
	case nil:
	}
	if len(decoded) < 1 {
		err = ErrInvalidFormat
		return
	}
	return decoded[1:], [2]byte{vByte, decoded[0]}, nil
}

// Base58Encode encodes in Base58Check with two version bytes.
func Base58Encode(data []byte, version [2]byte) string {
	buf := make([]byte, len(data)+1)
	buf[0] = version[1]
	copy(buf[1:], data)
	return base58.CheckEncode(buf, version[0])
}

func prfAddr(dst, ask []byte, t byte) {
	if len(dst) < 32 {
		panic("prfAddr called with a too small dst")
	}

	buf := make([]byte, 64)
	copy(buf, ask)
	buf[0] |= 0xc0
	buf[32] = t

	type compressor interface {
		SumNoPadding([]byte) []byte
	}

	d := sha256.New()
	d.Write(buf)
	d.(compressor).SumNoPadding(dst[:0])
}

func askToPKenc(ask []byte) []byte {
	var dst, src [32]byte
	prfAddr(src[:], ask, 1)
	curve25519.ScalarBaseMult(&dst, &src)
	return dst[:]
}

// KeyToAddress converts a raw spending key into a raw address.
func KeyToAddress(key []byte) ([]byte, error) {
	if len(key) != 32 || key[0]&0xf0 != 0 {
		return nil, ErrInvalidKey
	}

	addr := make([]byte, 64)
	prfAddr(addr, key, 0)
	copy(addr[32:], askToPKenc(key))
	return addr, nil
}

func KeyToViewingKey(key []byte) ([]byte, error) {
	if len(key) != 32 || key[0]&0xf0 != 0 {
		return nil, ErrInvalidKey
	}
	viewKey := make([]byte, 32)
	prfAddr(viewKey[:], key, 1)
	return viewKey, nil
}

// GenerateKey generates a new raw spending key.
func GenerateKey() []byte {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	b[0] &= 0x0f
	return b
}
