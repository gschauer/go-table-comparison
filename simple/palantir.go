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
	"strconv"
	"text/tabwriter"

	"github.com/palantir/pkg/tableprinter"
)

type addr struct {
	id    int
	name  string
	phone string
	email string
	qtty  int
}

var headers = []string{"#", "NAME", "PHONE", "EMAIL", "QTTY"}

func Palantir() string {
	b := new(bytes.Buffer)
	tw := tabwriter.NewWriter(b, 0, 0, 2, ' ', 0)

	cg := make(map[string]tableprinter.ColumnGetter)
	cg["#"] = func(row interface{}) string {
		return strconv.Itoa(row.(addr).id)
	}
	cg["NAME"] = func(row interface{}) string {
		return row.(addr).name
	}
	cg["PHONE"] = func(row interface{}) string {
		return row.(addr).phone
	}
	cg["EMAIL"] = func(row interface{}) string {
		return row.(addr).email
	}
	cg["QTTY"] = func(row interface{}) string {
		return strconv.Itoa(row.(addr).qtty)
	}

	var rows []interface{}
	for _, r := range data {
		rows = append(rows, addr{r[0].(int), r[1].(string), r[2].(string), r[3].(string), r[4].(int)})
	}

	t := tableprinter.New(tw, cg, false, true, true)
	if err := t.Print(headers, rows); err != nil {
		panic(err)
	}

	return b.String()
}
