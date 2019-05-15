package noscryptoclt

import (
	"strings"
)

// DecryptLoginPacket decrypts and returns the packet passed as argument
// using the client's login decryption algorithm.
func DecryptLoginPacket(packet string) (decryptedPacket string) {
	var buf strings.Builder

	for _, ch := range []byte(packet) {
		buf.WriteByte(ch - 0x0F)
	}

	decryptedPacket = buf.String()
	return decryptedPacket
}
