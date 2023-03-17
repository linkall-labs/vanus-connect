// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	cdkgo "github.com/vanus-labs/cdk-go"
<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/vanus-labs/connector/source/shopify/internal"
)

func main() {
	cdkgo.RunSource(internal.NewConfig, internal.NewShopifySource)
=======
	"github.com/vanus-labs/connector/source/http/internal"
)

func main() {
	cdkgo.RunSource(internal.NewConfig, internal.NewHTTPSource)
>>>>>>> d269259 (feat: add shopify source)
=======
	"github.com/vanus-labs/connector/source/shopify/internal"
)

func main() {
	cdkgo.RunSource(internal.NewConfig, internal.NewShopifySource)
>>>>>>> 2f93b62 (feat: add shopify source)
}
