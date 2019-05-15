package noscryptosrv

import (
	"strings"
)

// EncryptLoginPacket encrypts and returns the packet passed as argument
// using the server's login encryption algorithm.
func EncryptLoginPacket(packet string) (decryptedPacket string) {
	var buf strings.Builder

	for _, ch := range []byte(packet) {
		buf.WriteByte(uint8(ch + 15))
	}

	decryptedPacket = buf.String()
	return decryptedPacket
}