# Semver Version Compare

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/semver)](https://pkg.go.dev/github.com/go-zoox/semver)
[![Build Status](https://github.com/go-zoox/semver/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/semver/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/semver)](https://goreportcard.com/report/github.com/go-zoox/semver)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/semver/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/semver?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/semver.svg)](https://github.com/go-zoox/semver/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/semver.svg?label=Release)](https://github.com/go-zoox/semver/tags)

## Installation

To install the package, run:

```bash
go get github.com/go-zoox/semver
```

## Getting Started

```go
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
```

## License

GoZoox is released under the [MIT License](./LICENSE).
