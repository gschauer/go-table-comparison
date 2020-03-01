// Copyright 2020 The go-table-comparison authors
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

package simple

import (
	"bytes"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
)

func CheyneWallace() string {
	b := &bytes.Buffer{}
	t := tabby.NewCustom(tabwriter.NewWriter(b, 2, 0, 2, ' ', uint(0)))

	t.AddHeader("#", "NAME", "PHONE", "EMAIL", "QTTY")

	for _, r := range data {
		t.AddLine(r...)
	}

	t.Print()
	return b.String()
}
