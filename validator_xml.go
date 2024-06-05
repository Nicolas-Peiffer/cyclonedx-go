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
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
)

// ValidateXMLData validates the BOM in XML format against the specified schema version using encoding/xml package.
func ValidateXMLData(xmlData []byte, version SpecVersion) error {
	schemaPath := filepath.Join("schema", fmt.Sprintf("bom-%s.xsd", version.String()))
	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to open XML schema file: %v", err)
	}
	defer schemaFile.Close()

	// Read schema data
	var schemaData bytes.Buffer
	if _, err := io.Copy(&schemaData, schemaFile); err != nil {
		return fmt.Errorf("failed to read XML schema file: %v", err)
	}

	// Load the XSD schema and the XML catalog
	schema, err := xsd.Parse(schemaData.Bytes(), xsd.WithPath("schema/xmlcatalog.xml"))
	//schema, err := xsd.Parse(schemaData.Bytes())
	if err != nil {
		return fmt.Errorf("failed to parse XML schema: %v", err)
	}
	defer schema.Free()

	// Define struct for unmarshaling XML
	type BOMXML struct {
		// Define XML structure according to your schema
	}

	var bomXML BOMXML
	if err := xml.Unmarshal(xmlData, &bomXML); err != nil {
		return fmt.Errorf("failed to unmarshal XML data: %v", err)
	}

	// Validate the XML against the schema
	// Implement your validation logic here using bomXML and schemaData.Bytes()
	doc, err := libxml2.Parse(xmlData)
	if err != nil {
		return fmt.Errorf("failed to parse XML document: %v", err)
	}
	defer doc.Free()

	if err := schema.Validate(doc); err != nil {
		return fmt.Errorf("XML validation failed: %v", err)
	}

	return nil
}

// ValidateXML validates the BOM as XML.
func (bom *BOM) ValidateXML() error {
	version, err := getSpecVersionByXmlNamespace(bom.XMLNS)
	fmt.Print("ValidateXML version: ")
	fmt.Print(version.String())
	if err != nil {
		panic(err)
	}

	// encode BOM to JSON string
	var xmlData bytes.Buffer
	err = NewBOMEncoder(&xmlData, BOMFileFormatXML).
		SetPretty(false).
		Encode(bom)

	if err != nil {
		panic(err)
	}

	return ValidateXMLData(xmlData.Bytes(), version)
}
