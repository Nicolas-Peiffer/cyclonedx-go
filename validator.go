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

import "embed"

// Embed JSON schema and XML XSD schema in the binary
//
//go:embed schema/bom-1.2.schema.json
//go:embed schema/bom-1.3.schema.json
//go:embed schema/bom-1.4.schema.json
//go:embed schema/bom-1.5.schema.json
//go:embed schema/bom-1.6.schema.json
//go:embed schema/bom-1.0.xsd
//go:embed schema/bom-1.1.xsd
//go:embed schema/bom-1.2.xsd
//go:embed schema/bom-1.3.xsd
//go:embed schema/bom-1.4.xsd
//go:embed schema/bom-1.5.xsd
//go:embed schema/bom-1.6.xsd
//go:embed schema/spdx.xsd
var f embed.FS

// Validator interface describes BOM validator
type Validator interface {
	//Validate() (error, []error)
	Validate() error
}
