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
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

// ValidateJSONData validates the BOM in JSON format against the specified schema version.
func ValidateJSONData(jsonData []byte, version SpecVersion) error {
	// TODO version-1 because the SpecVersion start at 0 <-> 1.0 and
	schemaPath := filepath.Join("schema", fmt.Sprintf("bom-%s.schema.json", version.String()))
	schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", schemaPath))
	documentLoader := gojsonschema.NewBytesLoader(jsonData)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("failed to validate JSON: %v", err)
	}

	if !result.Valid() {
		var validationErrors string
		for _, desc := range result.Errors() {
			validationErrors += fmt.Sprintf("- %s\n", desc)
		}
		return fmt.Errorf("JSON validation failed: %s", validationErrors)
	}

	return nil
}

// ValidateJSON validates the BOM as JSON.
func (bom *BOM) ValidateJSON() error {
	fmt.Printf("bom.SpecVersion.String() %s\n", bom.SpecVersion.String())

	// encode BOM to JSON string
	var jsonData bytes.Buffer
	err := NewBOMEncoder(&jsonData, BOMFileFormatJSON).
		SetPretty(false).
		Encode(bom)

	if err != nil {
		panic(err)
	}

	return ValidateJSONData(jsonData.Bytes(), bom.SpecVersion)
}
