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

package main

import (
	"fmt"
	"strings"

	"github.com/gschauer/go-table-comparison/simple"
)

func main() {
	implByName := []struct {
		name string
		impl func() string
	}{
		{"github.com/alexeyco/simpletable", simple.AlexeyCo},
		{"github.com/bndr/gotabulate", simple.Bndr},
		{"github.com/caarlos0/tablewriter", simple.Caarlos0},
		{"github.com/cheynewallace/tabby", simple.CheyneWallace},
		{"github.com/gosuri/uitable", simple.GOsuri},
		{"github.com/jedib0t/go-pretty", simple.Jedib0t},
		{"github.com/olekukonko/tablewriter", simple.OlekuKonko},
		{"github.com/palantir/pkg/tableprinter", simple.Palantir},
		{"github.com/rodaine/table", simple.Rodaine},
		{"github.com/syohex/go-texttable", simple.Syohex},
		{"github.com/tatsushid/go-prettytable", simple.TatsushiD},
		{"text/tabwriter", simple.TabWriter},
	}

	for _, i := range implByName {
		fmt.Println(i.name)
		fmt.Println(i.impl())
		fmt.Println(strings.Repeat("=", 120))
	}
}
