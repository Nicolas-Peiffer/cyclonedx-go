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
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readXMLFromFileToBOMNoValidation(xmlFilePath string) (*BOM, error) {
	// Open the XML file
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		fmt.Printf("Error opening XML file: %v\n", err)
		return nil, err
	}
	defer xmlFile.Close()

	// Read the content of the file
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Printf("Error reading XML file: %v\n", err)
		return nil, err
	}

	// Unmarshal the XML data into a BOM struct
	var bom BOM
	if err := xml.Unmarshal(xmlData, &bom); err != nil {
		fmt.Printf("Error unmarshaling XML data: %v\n", err)
		return nil, err
	}

	// Now you have a BOM struct populated with the data from the XML file
	fmt.Printf("BOM: %+v\n", bom)
	return &bom, nil
}

func TestValidXMLBOM(t *testing.T) {
	bom, _ := readXMLFromFileToBOMNoValidation("testdata/valid-bom.xml")

	// Test JSON validation with a valid BOM
	assert.NoError(t, bom.ValidateXML(), "XML validation should succeed for a valid BOM")
}

func TestInvalidXMLBOM(t *testing.T) {
	invalidBOM, _ := readXMLFromFileToBOMNoValidation("testdata/invalid-crypto-protocols-cbom.xml")

	assert.Error(t, invalidBOM.ValidateXML(), "XML validation should fail for an invalid BOM")
}
