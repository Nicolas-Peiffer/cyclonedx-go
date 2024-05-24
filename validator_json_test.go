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
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readJSONFromFileToBOMNoValidation(jsonFilePath string) (*BOM, error) {
	// Read the JSON file
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Printf("Error opening JSON file: %v\n", err)
		return nil, err
	}
	defer jsonFile.Close()

	// Read the content of the file
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return nil, err
	}

	// Unmarshal the JSON data into a BOM struct
	var bom BOM
	if err := json.Unmarshal(jsonData, &bom); err != nil {
		fmt.Printf("Error unmarshaling JSON data: %v\n", err)
		return nil, err
	}

	// Now you have a BOM struct populated with the data from the JSON file
	fmt.Printf("BOM: %+v\n", bom)

	return &bom, nil
}

func TestValidJSONBOM(t *testing.T) {
	bom, _ := readJSONFromFileToBOMNoValidation("testdata/valid-bom.json")

	// Test JSON validation with a valid BOM
	assert.NoError(t, bom.ValidateJSON(), "JSON validation should succeed for a valid BOM")
}

func TestInvalidJSONBOM(t *testing.T) {
	// Test JSON validation with an invalid BOM
	// You can intentionally create an invalid BOM for testing
	// For example, set some required fields to invalid values
	invalidBOM, _ := readJSONFromFileToBOMNoValidation("testdata/invalid-crypto-protocols-cbom.json")

	fmt.Print("invalidBOM\n")
	fmt.Print(invalidBOM)

	fmt.Print(invalidBOM.ValidateJSON())

	assert.Error(t, invalidBOM.ValidateJSON(), "JSON validation should fail for an invalid BOM")
}
