package noscryptoclt

import (
	"bytes"
)

// EncryptGamePacket encrypts and returns the packet passed as argument
// using the session number provided and the client's game encryption
// algorithm.
func EncryptGamePacket(packet string, sessionNumber int) (encryptedPacket string) {
	encryptedPacket = firstEncryption(packet)
	encryptedPacket = secondEncryption(encryptedPacket, sessionNumber)

	return encryptedPacket
}

func firstEncryption(packet string) string {
	var encryptedPacket []byte
	var packetMask []bool
	var packetLength int

	var currentByte byte

	packetMask = generatePacketMask(packet)
	packetLength = len(packet)

	sequences := 0
	sequenceCounter := 0
	lastPosition := 0
	currentPosition := 0
	length := 0
	currentByte = 0
	for currentPosition <= packetLength {
		lastPosition = currentPosition
		for currentPosition < packetLength && packetMask[currentPosition] == false {
			currentPosition++
		}

		if currentPosition > 0 {
			length = currentPosition - lastPosition
			sequences = length / 0x7E
			for i := 0; i < length; i, lastPosition = i+1, lastPosition+1 {
				if i == sequenceCounter*0x7E {
					if sequences == 0 {
						encryptedPacket = append(encryptedPacket, byte(length-i))
					} else {
						encryptedPacket = append(encryptedPacket, 0x7E)
						sequences--
						sequenceCounter++
					}
				}

				encryptedPacket = append(encryptedPacket, packet[lastPosition]^0xFF)
			}
		}

		if currentPosition >= packetLength {
			break
		}

		lastPosition = currentPosition
		for currentPosition < packetLength && packetMask[currentPosition] == true {
			currentPosition++
		}

		if currentPosition > 0 {
			length = currentPosition - lastPosition
			sequences = length / 0x7E
			for i := 0; i < length; i, lastPosition = i+1, lastPosition+1 {
				if i == sequenceCounter*0x7E {
					if sequences == 0 {
						encryptedPacket = append(encryptedPacket, byte(length-i)|0x80)
					} else {
						encryptedPacket = append(encryptedPacket, 0x7E|0x80)
						sequences--
						sequenceCounter++
					}
				}

				currentByte = packet[lastPosition]
				switch currentByte {
				case 0x20:
					currentByte = 1
					break
				case 0x2D:
					currentByte = 2
					break
				case 0x2E:
					currentByte = 3
					break
				case 0xFF:
					currentByte = 0xE
					break
				default:
					currentByte -= 0x2C
					break
				}

				if currentByte != 0x00 {
					if i%2 == 0 {
						encryptedPacket = append(encryptedPacket, currentByte<<4)
					} else {
						encryptedPacket[len(encryptedPacket)-1] |= currentByte
					}
				}
			}
		}
	}

	encryptedPacket = append(encryptedPacket, 0xFF)

	return string(encryptedPacket)
}

func secondEncryption(packet string, sessionNumber int) string {
	var buf bytes.Buffer
	var sessionKey byte
	var xorKey byte

	//case 0
	sessionKey = byte(sessionNumber + 0x40)
	xorKey = 0x00

	switch (sessionNumber >> 6) & 0x03 {
	case 0:
		break
	case 1:
		sessionKey = -sessionKey
		break
	case 2:
		xorKey = 0xC3
		break
	case 3:
		sessionKey = -sessionKey
		xorKey = 0xC3
		break
	default:
		break
	}

	for _, ch := range []byte(packet) {
		buf.WriteByte((ch ^ xorKey) + sessionKey)
	}

	return buf.String()
}

func generatePacketMask(packet string) []bool {
	var mask []bool

	for _, ch := range []byte(packet) {
		ch := int(ch)

		if (ch & 0xFF) == 0x23 {
			mask = append(mask, false)
			continue
		}

		ch -= 0x20

		if (ch & 0xFF) == 0x00 {
			mask = append(mask, true)
			continue
		}

		ch += 0xF1

		if (ch & 0xFF) < 0x00 {
			mask = append(mask, true)
			continue
		}

		ch -= 0x0B

		if (ch & 0xFF) < 0x00 {
			mask = append(mask, true)
			continue
		}

		ch -= 0xC5

		if (ch & 0xFF) == 0x00 {
			mask = append(mask, true)
			continue
		}

		mask = append(mask, false)
	}

	return mask
}
