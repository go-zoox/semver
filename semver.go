package semver

import (
	"fmt"
	"strconv"
	"strings"
)

// IsEqual returns true if v1 and v2 are equal.
func IsEqual(v1, v2 string) bool {
	return v1 == v2
}

// IsGreatThan returns true if v1 is greater than v2.
func IsGreatThan(v1, v2 string) (bool, error) {
	if v1[0] == 'v' {
		v1 = v1[1:]
	}

	if v2[0] == 'v' {
		v2 = v2[1:]
	}

	v1s := strings.Split(v1, ".")
	v2s := strings.Split(v2, ".")

	if len(v1s) != 3 {
		return false, fmt.Errorf("version %s is not valid semver version", v1)
	}

	if len(v2s) != 3 {
		return false, fmt.Errorf("version %s is not valid semver version", v2)
	}

	v11, err := strconv.Atoi(v1s[0])
	if err != nil {
		return false, err
	}
	v12, err := strconv.Atoi(v1s[1])
	if err != nil {
		return false, err
	}
	v13, err := strconv.Atoi(v1s[2])
	if err != nil {
		return false, err
	}

	v21, err := strconv.Atoi(v2s[0])
	if err != nil {
		return false, err
	}
	v22, err := strconv.Atoi(v2s[1])
	if err != nil {
		return false, err
	}
	v23, err := strconv.Atoi(v2s[2])
	if err != nil {
		return false, err
	}

	if v11 < v21 {
		return false, nil
	} else if v11 > v21 {
		return true, nil
	}

	if v12 < v22 {
		return false, nil
	} else if v12 > v22 {
		return true, nil
	}

	if v13 <= v23 {
		return false, nil
	}

	return true, nil
}

// IsGreatThanOrEqual returns true if v1 is greater than or equal to v2.
func IsGreatThanOrEqual(v1, v2 string) (bool, error) {
	if IsEqual(v1, v2) {
		return true, nil
	}

	return IsGreatThan(v1, v2)
}

// IsLittleThan returns true if v1 is less than v2.
func IsLittleThan(v1, v2 string) (bool, error) {
	ok, err := IsGreatThanOrEqual(v1, v2)
	if err != nil {
		return false, err
	}

	return !ok, nil
}

// IsLittleThanOrEqual returns true if v1 is less than or equal to v2.
func IsLittleThanOrEqual(v1, v2 string) (bool, error) {
	if IsEqual(v1, v2) {
		return true, nil
	}

	return IsLittleThan(v1, v2)
}
