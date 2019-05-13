package noscryptoclt

import (
	"bytes"
)

// EncryptLoginPacket encrypts and returns the packet passed as argument
// using the client's login encryption algorithm.
func EncryptLoginPacket(packet string) (encryptedPacket string) {
	var buf bytes.Buffer

	for _, ch := range []byte(packet) {
		buf.WriteByte((ch ^ 0xC3) + 0x0F)
	}
	buf.WriteByte(0xD8) //"Ø"

	encryptedPacket = buf.String()
	return encryptedPacket
}
