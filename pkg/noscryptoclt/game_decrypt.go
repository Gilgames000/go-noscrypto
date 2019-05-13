package noscryptoclt

import (
	"bytes"
)

// DecryptGamePacket decrypts and returns the packet passed as argument
// using the session number provided and the client's game decryption
// algorithm. If multiple packets (separated by 0xFF) were passed as
// argument, they will all be decrypted and separated by '\n' in the
// returned string.
func DecryptGamePacket(packet string) (decryptedPacket string) {
	var result bytes.Buffer
	var currentPacket string
	var currentByte byte
	var firstChar byte
	var secondChar byte

	charTable := []byte{
		' ', '-', '.', '0', '1', '2', '3',
		'4', '5', '6', '7', '8', '9', 'n'}
	packetBytes := []byte(packet)
	substringLength := 0
	i := 0

	for i < len(packetBytes) {
		currentByte = packetBytes[i]
		i++

		// The 0xFF byte represents the end of the current packet
		if currentByte == 0xFF {
			result.WriteString(currentPacket)
			if i != len(packetBytes) {
				result.WriteByte('\n')
			}
			currentPacket = ""
			continue
		}

		// The original packe was encrypted by dividing the whole string in
		// substrings long 0x7E bytes each
		substringLength = int(currentByte & 0x7F)

		if currentByte&0x80 != 0x00 {
			// Bytes below 0x80 need to be decrypted using the character table
			// defined above
			for substringLength != 0 {
				if i < len(packetBytes) {
					currentByte = packetBytes[i]
					i++

					firstChar = charTable[((currentByte&0xF0)>>4)-1]
					if firstChar != 0x6E {
						currentPacket += string(firstChar)
					}

					if substringLength <= 1 {
						break
					}

					secondChar = charTable[(currentByte&0xF)-1]
					if secondChar != 0x6E {
						currentPacket += string(secondChar)
					}

					substringLength -= 2
				} else {
					substringLength--
				}
			}
		} else {
			// All other bytes can be decrypted by inverting their
			// bits (ones' complement)
			for substringLength != 0 {
				if i < len(packetBytes) {
					currentPacket += string(packetBytes[i] ^ 0xFF)
					i++
				}
				substringLength--
			}
		}
	}

	decryptedPacket = result.String()
	return decryptedPacket
}
