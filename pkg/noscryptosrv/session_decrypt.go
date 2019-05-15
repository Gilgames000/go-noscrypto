package noscryptosrv

import (
	"strings"
)

// DecryptSessionPacket decrypts and returns the packet passed as argument
// using the server's session decryption algorithm.
func DecryptSessionPacket(packet string) (decryptedPacket string) {
	var firstbyte byte
	var secondbyte byte
	var result strings.Builder

	packetBytes := []byte(packet)

	for i := 1; i < len(packetBytes); i++ {
		if packetBytes[i] == 0x0E {
			break
		}

		firstbyte = packetBytes[i] - 0xF
		secondbyte = firstbyte

		secondbyte &= 0xF0
		firstbyte = firstbyte - secondbyte
		secondbyte >>= 0x4

		evalByte := func(b byte) byte {
			switch b {
			case 0:
			case 1:
				b = 0x20 // " "
				break
			case 2:
				b = 0x2D // "-"
				break
			case 3:
				b = 0x2E // "."
				break
			default:
				b += 0x2C
				break
			}

			return b
		}

		result.WriteByte(evalByte(secondbyte))
		result.WriteByte(evalByte(firstbyte))
	}

	decryptedPacket = result.String()
	return decryptedPacket
}
