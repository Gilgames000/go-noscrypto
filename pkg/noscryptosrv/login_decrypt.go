package noscryptosrv

import (
	"bytes"
)

// DecryptLoginPacket decrypts and returns the packet passed as argument
// using the server's login decryption algorithm.
func DecryptLoginPacket(packet string) (decryptedPacket string) {
	var buf bytes.Buffer

	for _, ch := range []byte(packet[:len(packet)-2]) {
		buf.WriteByte((ch - 0x0F) ^ 0xC3)
	}

	decryptedPacket = buf.String()
	return decryptedPacket
}
