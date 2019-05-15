package noscryptosrv

import (
	"github.com/gilgames000/go-noscrypto/pkg/noscryptoclt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Decrypt(t *testing.T) {
	packet := "testing packet"
	encrypted := noscryptoclt.EncryptLoginPacket(packet)
	decrypted := DecryptLoginPacket(encrypted)

	assert.Equal(t, packet, decrypted, "Decrypted packet should be the same as the original one")
}