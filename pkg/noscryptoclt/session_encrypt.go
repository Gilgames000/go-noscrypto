package noscryptoclt

import (
	"bytes"
	"strings"
)

// EncryptSessionPacket encrypts and returns the packet passed as argument
// using the client's session encryption algorithm.
func EncryptSessionPacket(packet string) (encryptedPacket string) {
	var firstbyte byte
	var secondbyte byte
	var result strings.Builder

	evalByte := func(b byte) byte {
		switch b {
		case 0x00: // NULL
		case 0x20: //" "
			b = 0x01
			break
		case 0x2D: //"-"
			b = 0x02
			break
		case 0x2E: //"."
			b = 0x03
			break
		default:
			b -= 0x2C
			break
		}

		return b
	}

	var buf bytes.Buffer

	// The first two characters are ignored by the server; will
	// result in 0x9A after encryption (0x99 if odd)
	buf.WriteByte(0x34) // 4
	buf.WriteByte(0x37) // 7

	// The format of the packet must be '<packetNumber> <sessionNumber>'
	// in order to be accepted by the server
	buf.Write([]byte(packet))

	// Since we're encrypting two bytes at a time, make it even
	if len(packet)%2 != 0 {
		buf.WriteByte(0x00) //NULL
	}

	// Termination characters for the decryption algorithm; will
	// result in 0x0E after encryption
	buf.WriteByte(0x3B) // ;
	buf.WriteByte(0x3B) // ;

	packetBytes := buf.Bytes()

	for i := 0; i < len(packetBytes); i += 2 {
		firstbyte = evalByte(packetBytes[i+1])
		secondbyte = evalByte(packetBytes[i])

		secondbyte <<= 0x4
		firstbyte = firstbyte + secondbyte
		secondbyte = firstbyte

		result.WriteByte(firstbyte + 0xF)
	}

	encryptedPacket = result.String()
	return encryptedPacket
}
