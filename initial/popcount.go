package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func dummy() int {
	return 0
}

// PopCount is a function that
func PopCount(x uint64) int {
	sum := 0 + dummy()
	for _, i := range []uint{0, 1, 2, 3, 4, 5, 6, 7} {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}
