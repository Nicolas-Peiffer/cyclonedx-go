package cyclonedx

import (
	"fmt"
	"testing"

	"github.com/blang/semver/v4"
	"github.com/stretchr/testify/assert"
)

func TestSpecVersion_String(t *testing.T) {
	assert := assert.New(t)

	fmt.Printf("SpecVersion1_0.String(): %s", SpecVersion1_0.String())

	assert.Equal("1.0.0", SpecVersion1_0.String(), "SpecVersion1_0 should be '1.0.0'")
	assert.Equal("1.1.0", SpecVersion1_1.String(), "SpecVersion1_1 should be '1.1.0'")
	assert.Equal("1.2.0", SpecVersion1_2.String(), "SpecVersion1_2 should be '1.2.0'")
	assert.Equal("1.3.0", SpecVersion1_3.String(), "SpecVersion1_3 should be '1.3.0'")
	assert.Equal("1.4.0", SpecVersion1_4.String(), "SpecVersion1_4 should be '1.4.0'")
	assert.Equal("1.5.0", SpecVersion1_5.String(), "SpecVersion1_5 should be '1.5.0'")
	assert.Equal("1.6.0", SpecVersion1_6.String(), "SpecVersion1_6 should be '1.6.0'")
}

func TestSpecVersion_Compare(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(-1, SpecVersion1_0.Compare(SpecVersion1_1), "SpecVersion1_0 should be less than SpecVersion1_1")
	assert.Equal(1, SpecVersion1_2.Compare(SpecVersion1_1), "SpecVersion1_2 should be greater than SpecVersion1_1")
	assert.Equal(0, SpecVersion1_3.Compare(SpecVersion1_3), "SpecVersion1_3 should be equal to SpecVersion1_3")
	assert.Equal(-1, SpecVersion1_4.Compare(SpecVersion1_5), "SpecVersion1_4 should be less than SpecVersion1_5")
	assert.Equal(1, SpecVersion1_6.Compare(SpecVersion1_5), "SpecVersion1_6 should be greater than SpecVersion1_5")
}

func TestSpecVersion_InvalidVersion(t *testing.T) {
	assert := assert.New(t)

	// Testing invalid version string should panic
	assert.Panics(func() {
		SpecVersion(semver.MustParse("invalid.version"))
	}, "Invalid version string should cause a panic")
}
