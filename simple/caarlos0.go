// Copyright 2022 The go-table-comparison authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package simple

import (
	"strconv"
	"strings"

	"github.com/caarlos0/tablewriter"
)

func Caarlos0() string {
	b := &strings.Builder{}
	err := tablewriter.Render(b, data, []string{"#", "NAME", "PHONE", "EMAIL", "QTTY"},
		func(row []any) ([]string, error) {
			return []string{
				strconv.Itoa(row[0].(int)),
				row[1].(string),
				row[2].(string),
				row[3].(string),
				strconv.Itoa(row[4].(int))}, nil
		})

	if err != nil {
		panic(err)
	}

	return b.String()
}
