package semver

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestIsGreatThan(t *testing.T) {
	res, _ := IsGreatThan("v2.0.0", "v1.0.0")
	testify.Assert(t, res, "v2.0.0 should be greater than v1.0.0")

	res, _ = IsGreatThan("v2.1.0", "v2.0.0")
	testify.Assert(t, res, "v2.1.0 should not be greater than v2.0.0")

	res, _ = IsGreatThan("v2.0.1", "v2.0.0")
	testify.Assert(t, res, "v2.0.1 should not be greater than v2.0.0")

	res, _ = IsGreatThan("v2.0.0", "v2.0.0")
	testify.Assert(t, !res, "v2.0.0 should not be greater than v2.0.0")

	res, _ = IsGreatThan("v1.0.0", "v2.0.0")
	testify.Assert(t, !res, "v1.0.0 should not be greater than v2.0.0")
}

func TestIsGreatThanOrEqual(t *testing.T) {
	res, _ := IsGreatThanOrEqual("v2.0.0", "v1.0.0")
	testify.Assert(t, res, "v2.0.0 should be greater than v1.0.0")

	res, _ = IsGreatThanOrEqual("v2.1.0", "v2.0.0")
	testify.Assert(t, res, "v2.1.0 should not be greater than v2.0.0")

	res, _ = IsGreatThanOrEqual("v2.0.1", "v2.0.0")
	testify.Assert(t, res, "v2.0.1 should not be greater than v2.0.0")

	res, _ = IsGreatThanOrEqual("v2.0.0", "v2.0.0")
	testify.Assert(t, res, "v2.0.0 should not be greater than v2.0.0")

	res, _ = IsGreatThanOrEqual("v1.0.0", "v2.0.0")
	testify.Assert(t, !res, "v1.0.0 should not be greater than v2.0.0")
}

func TestIsLittleThan(t *testing.T) {
	res, _ := IsLittleThan("v2.0.0", "v1.0.0")
	testify.Assert(t, !res, "v2.0.0 should be greater than v1.0.0")

	res, _ = IsLittleThan("v2.1.0", "v2.0.0")
	testify.Assert(t, !res, "v2.1.0 should not be greater than v2.0.0")

	res, _ = IsLittleThan("v2.0.1", "v2.0.0")
	testify.Assert(t, !res, "v2.0.1 should not be greater than v2.0.0")

	res, _ = IsLittleThan("v2.0.0", "v2.0.0")
	testify.Assert(t, !res, "v2.0.0 should not be greater than v2.0.0")

	res, _ = IsLittleThan("v1.0.0", "v2.0.0")
	testify.Assert(t, res, "v1.0.0 should not be greater than v2.0.0")
}

func TestIsLittleThanOrEqual(t *testing.T) {
	res, _ := IsLittleThanOrEqual("v2.0.0", "v1.0.0")
	testify.Assert(t, !res, "v2.0.0 should be greater than v1.0.0")

	res, _ = IsLittleThanOrEqual("v2.1.0", "v2.0.0")
	testify.Assert(t, !res, "v2.1.0 should not be greater than v2.0.0")

	res, _ = IsLittleThanOrEqual("v2.0.1", "v2.0.0")
	testify.Assert(t, !res, "v2.0.1 should not be greater than v2.0.0")

	res, _ = IsLittleThanOrEqual("v2.0.0", "v2.0.0")
	testify.Assert(t, res, "v2.0.0 should not be greater than v2.0.0")

	res, _ = IsLittleThanOrEqual("v1.0.0", "v2.0.0")
	testify.Assert(t, res, "v1.0.0 should not be greater than v2.0.0")
}
