package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	input := "bgvyzdsv"
	number := 0
	solution := 0
	prefix := "00000"
	for {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", input, number)))
		hashString := hex.EncodeToString(hash[:])
		fmt.Printf("\rhashString: %v", hashString)
		if strings.HasPrefix(hashString, prefix) {
			solution = number
			break
		}
		number++
	}
	fmt.Printf("solution: %v\n", solution)
}
