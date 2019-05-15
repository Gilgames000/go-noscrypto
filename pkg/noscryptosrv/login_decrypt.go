package noscryptosrv

import (
	"strings"
)

// DecryptLoginPacket decrypts and returns the packet passed as argument
// using the server's login decryption algorithm.
func DecryptLoginPacket(packet string) (decryptedPacket string) {
	var buf strings.Builder

	for _, ch := range []byte(packet[:len(packet)-1]) {
		buf.WriteByte((ch - 0x0F) ^ 0xC3)
	}

	decryptedPacket = buf.String()
	return decryptedPacket
}
