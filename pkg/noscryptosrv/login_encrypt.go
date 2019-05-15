package noscryptosrv

import (
	"bytes"
)

// EncryptLoginPacket encrypts and returns the packet passed as argument
// using the server's login encryption algorithm.
func EncryptLoginPacket(packet string) (decryptedPacket string) {
	var buf bytes.Buffer

	for _, ch := range []byte(packet + " ") {
		buf.WriteByte(uint8(ch + 15))
	}

	bytesBuffer := buf.Bytes()
	bytesBuffer[buf.Len()-1] = 25 // ^Y

	decryptedPacket = string(bytesBuffer)
	return decryptedPacket
}