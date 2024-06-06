// This file is part of CycloneDX Go
//
// Licensed under the Apache License, Version 2.0 (the “License”);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an “AS IS” BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0
// Copyright (c) OWASP Foundation. All Rights Reserved.

package cyclonedx

import (
	"errors"
	"fmt"
	"strings"

	"github.com/blang/semver/v4"
)

// Alias semver.Version as SpecVersion
type SpecVersion semver.Version

// Method to get the version string
func (v SpecVersion) String() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

// Method to compare SpecVersion with another SpecVersion
func (v SpecVersion) Compare(other SpecVersion) int {
	return semver.Version(v).Compare(semver.Version(other))
}

// Return true if semantic version v is Greater Or Equal than semantic version
// o.
func (v SpecVersion) IsGreaterOrEqualVersion(o SpecVersion) bool {
	// Compare compares Versions v to o:
	// -1 == v is less than o
	// 0 == v is equal to o
	// 1 == v is greater than o
	comparisonResult := v.Compare(o)

	switch comparisonResult {
	case 0, 1:
		return true
	default:
		return false
	}
}

var (
	SpecVersion1_0 = SpecVersion(semver.MustParse("1.0.0")) // 1.0
	SpecVersion1_1 = SpecVersion(semver.MustParse("1.1.0")) // 1.1
	SpecVersion1_2 = SpecVersion(semver.MustParse("1.2.0")) // 1.2
	SpecVersion1_3 = SpecVersion(semver.MustParse("1.3.0")) // 1.3
	SpecVersion1_4 = SpecVersion(semver.MustParse("1.4.0")) // 1.4
	SpecVersion1_5 = SpecVersion(semver.MustParse("1.5.0")) // 1.5
	SpecVersion1_6 = SpecVersion(semver.MustParse("1.6.0")) // 1.6
)

var ErrInvalidSpecVersion = errors.New("invalid specification version")

// TODO: this should use XML Catalogs instead. See github.com/jteeuwen/go-pkg-xmlx
var xmlNamespaces = map[string]string{
	SpecVersion1_0.String(): "http://cyclonedx.org/schema/bom/1.0",
	SpecVersion1_1.String(): "http://cyclonedx.org/schema/bom/1.1",
	SpecVersion1_2.String(): "http://cyclonedx.org/schema/bom/1.2",
	SpecVersion1_3.String(): "http://cyclonedx.org/schema/bom/1.3",
	SpecVersion1_4.String(): "http://cyclonedx.org/schema/bom/1.4",
	SpecVersion1_5.String(): "http://cyclonedx.org/schema/bom/1.5",
	SpecVersion1_6.String(): "http://cyclonedx.org/schema/bom/1.6",
}

func getSpecVersionByXmlNamespace(value string) (SpecVersion, error) {
	for k, v := range xmlNamespaces {
		if v == value {
			res, _ := SafeParseSemVer(k)
			return SpecVersion(res), nil
		}
	}
	return SpecVersion(semver.MustParse("0.0.0")), errors.New("value not found in xmlNamespaces")
}

// SafeParseSemVer is a method that parses a version string that may be in
// major.minor or major.minor.patch format. If major.minor, then add an extra
// zero as patch: major.minor.0.
func SafeParseSemVer(versionString string) (SpecVersion, error) {
	// Preprocess the version string to add ".0" if it's missing the patch version
	if strings.Count(versionString, ".") == 1 {
		versionString += ".0"
	}

	// Parse the version string
	version, err := semver.Parse(versionString)
	if err != nil {
		return SpecVersion{}, fmt.Errorf("error parsing version: %v", err)
	}

	return SpecVersion(version), nil
}
