package testHelpers

import "errors"

func CheckStringSliceEqual(sl1, sl2 []string) error {
	if len(sl1) != len(sl2) {
		return errors.New("slices have != lengths")
	}

	for _, v1 := range sl1 {
		found := false
		for _, v2 := range sl2 {
			if v1 == v2 {
				found = true
				break
			}
		}
		if !found {
			return errors.New(v1 + " has not been found in slice 2")
		}
	}
	return nil
}
