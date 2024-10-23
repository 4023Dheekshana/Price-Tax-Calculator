package conversion

import (
	"errors"
	"strconv"
)

func StringToFloats(lines []string) ([]float64, error) {
	floatPrices := []float64{}
	for _, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return nil, errors.New("converting string to float is failed")
		}
		floatPrices = append(floatPrices, floatPrice)

	}
	return floatPrices, nil
}
