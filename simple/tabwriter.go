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
	"fmt"
	"text/tabwriter"
)

func TabWriter() string {
	b := &bytes.Buffer{}
	w := tabwriter.NewWriter(b, 2, 0, 2, ' ', 0)

	_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", "#", "NAME", "PHONE", "EMAIL", "QTTY")
	_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", "----", "----", "----", "----", "----")

	for _, r := range data {
		_, _ = fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%d\n", r[0], r[1], r[2], r[3], r[4])
	}

	_ = w.Flush()
	return b.String()
}
