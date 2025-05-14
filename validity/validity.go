package validity

import (
	"errors"
)

var (
	list []rune
)

func init() {
	list = []rune{'e', 'd'}
}
func Validation(value string, key int) (string, error) {

	switch key {
	case 1:
		if len(value) < 10 {
			return value, errors.New("the provided key is invalid. Please check and try again")
		}
	case 2:
		for _, v := range list {
			if string(v) == value {
				return value, nil
			}
		}
		return value, errors.New("please choose an operation: (e)encrypt / (d)decrypt")

	default:
		return value, nil

	}
	return value, nil

}
