# go-noscrypto
Golang implementation of the packet encryption/decryption algorithms used by the
MMORPG NosTale.

## Installation
```
go get -u github.com/gilgames000/go-noscrypto/
```

## Example
```go
package main

import (
	"fmt"
	"github.com/gilgames000/go-noscrypto/pkg/noscryptoclt"
)

func main() {
	sessionNumber := 10
	encryptedPacket := noscryptoclt.EncryptGamePacket("say 1 1337 0 Hello World", sessionNumber)
	fmt.Println("encryptedPacket : ", encryptedPacket)
}
```

## Packages
- [noscryptoclt](https://github.com/Gilgames000/go-noscrypto/tree/master/pkg/noscryptoclt) - algorithms used by NosTale Gameforge client
- [noscryptosrv](https://github.com/Gilgames000/go-noscrypto/tree/master/pkg/noscryptosrv) - algorithms used by NosTale Gameforge server (incomplete)

## Documentation
You can check the documentation online on [godoc.org](https://godoc.org/?q=go-noscrypto)

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/Gilgames000/go-noscrypto/blob/master/LICENSE) file for details

## Acknowledgements
- [All NosTale Cryptography](http://www.bordergame.it/Thread-All-Nostale-Cryptography)
