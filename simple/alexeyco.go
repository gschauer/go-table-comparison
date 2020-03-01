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
	"strconv"

	"github.com/alexeyco/simpletable"
)

func AlexeyCo() string {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Text: "#"},
			{Text: "NAME"},
			{Text: "PHONE"},
			{Text: "EMAIL"},
			{Text: "QTTY"},
		},
	}

	for _, r := range data {
		cs := []*simpletable.Cell{
			{Text: strconv.Itoa(r[0].(int))},
			{Text: r[1].(string)},
			{Text: r[2].(string)},
			{Text: r[3].(string)},
			{Text: strconv.Itoa(r[4].(int))},
		}

		table.Body.Cells = append(table.Body.Cells, cs)
	}

	return table.String()
}
