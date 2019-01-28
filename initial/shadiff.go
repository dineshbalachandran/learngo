package main

import (
	"crypto/sha256"
	"fmt"
)

func shadiff() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	sum := 0
	for i := 0; i < 32; i++ {
		d := byte(c1[i] ^ c2[i])
		for j := uint(0); j < 8; j++ {
			sum += int((d >> j) & 1)
		}
	}

	fmt.Printf("%x\n%x\n%d\n", c1, c2, sum)
}
