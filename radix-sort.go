package main

import (
	"fmt"
	"encoding/binary"
	"bytes"
)

const bytesInDigit = 4 // 32bit / 8bit 
const bitsInByte = 256
const maxbit = -1 << 31

func countingSort(bytesArray [][]byte) [][]byte {
	// [] [] [] [] == digit
	// outer loop : for each byte

	counts := make([][][]byte, bitsInByte)

	for i := 0; i < bytesInDigit; i++ {
		for _, numberBytes := range bytesArray {
			counts[numberBytes[i]] = append(counts[numberBytes[i]], numberBytes)
		}
	
		j := 0
		// counts are now sorted by little bytes 
		for k, numBytes := range counts {
			copy(bytesArray[j:], numBytes)
			j += len(numBytes) // if there was no number, j is still the same
			counts[k] = numBytes[:0] // empty the number in counts
		}
	}

	return bytesArray
}


func sort(array []int32) []int32 {
	bytesArray := make([][]byte, len(array))
	buf := bytes.NewBuffer(nil)
	for i, num := range array {
		// little endian: read small bits first for the purpose of counting sort
		binary.Write(buf, binary.LittleEndian, num^maxbit)
		bytesArray[i] = make([]byte, bytesInDigit)
		buf.Read(bytesArray[i])
	}
	bytesArray = countingSort(bytesArray)

	var num int32
	for i, numBytes := range bytesArray {
		buf.Write(numBytes)
		binary.Read(buf, binary.LittleEndian, &num)
		array[i] = num^maxbit
	}
	return array
}


func run() {
	a := []int32{134, 5, 34, 244, 21, -10, 0,  47}
	fmt.Println(a)
	a = sort(a)
	fmt.Println(a)
}

func main() {
	run()
}
