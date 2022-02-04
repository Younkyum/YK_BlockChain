package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	difficutly := 6
	target := strings.Repeat("0", difficutly)
	nounce := 1
	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello"+fmt.Sprint(nounce))))
		fmt.Printf("Hash:%s\nTarget:%s\nNounce:%d\n\n", hash, target, nounce)
		if strings.HasPrefix(hash, target) {
			return
		} else {
			nounce++
		}
	}
}
