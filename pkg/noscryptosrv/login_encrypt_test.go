package noscryptosrv

import (
	"github.com/gilgames000/go-noscrypto/pkg/noscryptoclt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EncryptLoginPacket(t *testing.T) {
	packet := "testing packet"
	encrypted := EncryptLoginPacket(packet)
	decrypted := noscryptoclt.DecryptLoginPacket(encrypted)

	assert.Equal(t, packet, decrypted, "Decrypted packet should be the same as the one encrypted by client encryption algorithm")
}
