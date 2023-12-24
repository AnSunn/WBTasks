package task8

import "fmt"

func setBit(value int64, bitIndex int, bitValue bool) int64 {
	// Mask to set or reset bit
	mask := int64(1) << uint(bitIndex)

	if bitValue {
		//Set 1
		return value | mask
	} else {
		// Reset 0
		return value &^ mask
	}
}

func Launch() {
	var number int64
	var bitIndex int
	var bitValue bool

	// Enter data
	fmt.Print("Enter value int64: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}

	fmt.Print("Enter bit index: ")
	_, err = fmt.Scan(&bitIndex)
	if err != nil {
		return
	}

	fmt.Print("Enter bit value (1/0): ")
	_, err = fmt.Scan(&bitValue)
	if err != nil {
		return
	}
	// Set or reset bit
	result := setBit(number, bitIndex, bitValue)
	// Display results
	fmt.Printf("The result is: %d\n", result)
}
