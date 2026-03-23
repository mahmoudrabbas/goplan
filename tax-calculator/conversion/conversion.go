package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for i, val := range strings {
		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, errors.New("Error Converting String into Float.")
		}

		floats[i] = floatVal
	}

	return floats, nil
}
