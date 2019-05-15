package noscryptoclt

import (
	"strings"
)

// EncryptLoginPacket encrypts and returns the packet passed as argument
// using the client's login encryption algorithm.
func EncryptLoginPacket(packet string) (encryptedPacket string) {
	var buf strings.Builder

	for _, ch := range []byte(packet) {
		buf.WriteByte((ch ^ 0xC3) + 0x0F)
	}
	buf.WriteByte(0xD8) //"Ø"

	encryptedPacket = buf.String()
	return encryptedPacket
}
