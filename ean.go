package brdoc

import (
	"fmt"
	"strconv"
)

func IsEanValid(ean string) bool {
	return ValidEan8(ean) || ValidEan13(ean) || ValidEan12(ean) || ValidEan14(ean)
}

func ValidEan8(ean string) bool {
	return validCode(ean, 8)
}

func ValidEan12(upc string) bool {
	return validCode(upc, 12)

}

func ValidEan13(ean string) bool {
	return validCode(ean, 13)
}

func ValidEan14(ean string) bool {
	return validCode(ean, 14)

}

func validCode(ean string, size int) bool {
	checksum, err := checksum(ean, size)

	return err == nil && strconv.Itoa(checksum) == ean[size-1:size]
}

func ChecksumEan8(ean string) (int, error) {
	return checksum(ean, 8)
}

func ChecksumEan12(upc string) (int, error) {
	return checksum(upc, 12)
}

func ChecksumEan13(ean string) (int, error) {
	return checksum(ean, 13)
}

func ChecksumEan14(ean string) (int, error) {
	return checksum(ean, 14)
}


//GTIN-8  (N8)
//GTIN-12 (N12) UPC
//GTIN-13 (N13)
//GTIN-14 (N14)
//GSIN    (N17)
//SSCC    (N18)
//Step 1: Multiply value of each position by
//x3x1
//Step 2: Add results together to create sum
//Step 3: Subtract the sum from nearest equal or higher multiple of ten = Check Digit
func checksum(ean string, size int) (int, error) {
	if len(ean) != size {
		return -1, fmt.Errorf("incorrect ean %v to compute a checksum", ean)
	}

	code := ean[:size-1]
	multiplyWhenEven := size%2 == 0
	sum := 0

	for i, v := range code {
		value, err := strconv.Atoi(string(v))

		if err != nil {
			return -1, fmt.Errorf("contains non-digit: %q", v)
		}

		if (i%2 == 0) == multiplyWhenEven {
			sum += 3 * value
		} else {
			sum += value
		}
	}

	return (10 - sum%10) % 10, nil
}



