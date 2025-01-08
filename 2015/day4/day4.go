package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	input := "bgvyzdsv"
	number := 0
	for {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", input, number)))
		hashString := hex.EncodeToString(hash[:])
		if hashString[0:6] == "000000" {
			fmt.Printf("Solution: %v %s\n", number, hashString)
			break
		}
		number++
	}
}
