# Go Table Library Comparison
There are several libraries for rending tables in console applications written Go.
However, it's difficult to compare their features as well as their performance.

Below you can find an subjective overview of several table rendering libraries:
* [golang.org/pkg/text/tabwriter][tabwriter]
* [github.com/jedib0t/go-pretty][go-pretty]
* [github.com/tatsushid/go-prettytable][go-prettytable]
* [github.com/bndr/gotabulate][gotabulate]
* [github.com/syohex/go-texttable][go-texttable]
* [github.com/alexeyco/simpletable][simpletable]
* [github.com/cheynewallace/tabby][tabby]
* [github.com/rodaine/table][table]
* [github.com/palantir/pkg][tableprinter]
* [github.com/olekukonko/tablewriter][tablewriter]
* [github.com/gosuri/uitable][uitable]

## Feature Comparison
Please note that the absence of a feature does not mean that there is no support at all.

For example, many libraries don't have color support built-in but are compatible with common ANSI color libraries.

| Feature            |     [tabwriter]      |    [go-pretty]     |  [go-prettytable]   |     [gotabulate]     |   [go-texttable]   |    [simpletable]     |       [tabby]        |       [table]        |    [tableprinter]    |    [tablewriter]     |      [uitable]       |
| :----------------- | :------------------: | :----------------: | :-----------------: | :------------------: | :----------------: | :------------------: | :------------------: | :------------------: | :------------------: | :------------------: | :------------------: |
| Github last commit |                      |   ![c_go-pretty]   | ![c_go-prettytable] |   ![c_gotabulate]    | ![c_go-texttable]  |   ![c_simpletable]   |      ![c_tabby]      |      ![c_table]      |  ![c_tableprinter]   |   ![c_tablewriter]   |     ![c_uitable]     |
| Github popularity  |  language built-in   |   ![s_go-pretty]   | ![s_go-prettytable] |   ![s_gotabulate]    | ![s_go-texttable]  |   ![s_simpletable]   |      ![s_tabby]      |      ![s_table]      |  ![s_tableprinter]   |   ![s_tablewriter]   |     ![s_uitable]     |
|                    |                      |                    |                     |                      |                    |                      |                      |                      |                      |                      |                      |
| performance        |  :heavy_check_mark:  | :heavy_check_mark: | :heavy_check_mark:  |  :heavy_check_mark:  | :heavy_check_mark: | :small_blue_diamond: |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  | :small_blue_diamond: |         :x:          |
|                    |                      |                    |                     |                      |                    |                      |                      |                      |                      |                      |                      |
| built-in colors    |  :heavy_check_mark:  | :heavy_check_mark: |         :x:         |         :x:          |        :x:         |         :x:          |         :x:          |         :x:          |         :x:          |  :heavy_check_mark:  |         :x:          |
| predefined styles  |         :x:          | :heavy_check_mark: |         :x:         |  :heavy_check_mark:  |        :x:         |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          |         :x:          |         :x:          |
| custom styles      |         :x:          | :heavy_check_mark: | :heavy_check_mark:  |         :x:          |        :x:         |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          |  :heavy_check_mark:  |         :x:          |
|                    |                      |                    |                     |                      |                    |                      |                      |                      |                      |                      |                      |
| align text         | :small_blue_diamond: | :heavy_check_mark: | :heavy_check_mark:  | :small_blue_diamond: |        :x:         |  :heavy_check_mark:  | :small_blue_diamond: |         :x:          | :small_blue_diamond: |  :heavy_check_mark:  | :small_blue_diamond: |
| borders            | :small_blue_diamond: | :heavy_check_mark: |         :x:         |  :heavy_check_mark:  |        :x:         |  :heavy_check_mark:  | :small_blue_diamond: |         :x:          | :small_blue_diamond: |  :heavy_check_mark:  |         :x:          |
| cell span          |         :x:          | :heavy_check_mark: |         :x:         |         :x:          |        :x:         |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          |  :heavy_check_mark:  |         :x:          |
| custom renderer    |         :x:          | :heavy_check_mark: |         :x:         |         :x:          |        :x:         |  :heavy_check_mark:  |         :x:          | :small_blue_diamond: |         :x:          |  :heavy_check_mark:  |         :x:          |
| footer             |         :x:          | :heavy_check_mark: |         :x:         |         :x:          |        :x:         |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          |  :heavy_check_mark:  |         :x:          |
| multiline          |         :x:          | :heavy_check_mark: |         :x:         |  :heavy_check_mark:  |        :x:         |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          |  :heavy_check_mark:  |         :x:          |
|                    |                      |                    |                     |                      |                    |                      |                      |                      |                      |                      |                      |
| string values      |  :heavy_check_mark:  | :heavy_check_mark: | :heavy_check_mark:  |  :heavy_check_mark:  | :heavy_check_mark: |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |
| interface values   |  :heavy_check_mark:  | :heavy_check_mark: | :heavy_check_mark:  |  :heavy_check_mark:  |        :x:         |         :x:          |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |         :x:          |  :heavy_check_mark:  |
| arbitrary structs  |         :x:          |        :x:         |         :x:         |         :x:          |        :x:         |         :x:          |         :x:          |         :x:          |  :heavy_check_mark:  |         :x:          |         :x:          |
|                    |                      |                    |                     |                      |                    |                      |                      |                      |                      |                      |                      |
| ASCII output       |  :heavy_check_mark:  | :heavy_check_mark: | :heavy_check_mark:  |  :heavy_check_mark:  | :heavy_check_mark: |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |
| CSV output         |         :x:          | :heavy_check_mark: |         :x:         |         :x:          |        :x:         |         :x:          |         :x:          |         :x:          |         :x:          |         :x:          |         :x:          |
| HTML output        |         :x:          | :heavy_check_mark: |         :x:         |         :x:          |        :x:         |         :x:          |         :x:          |         :x:          |         :x:          |         :x:          |         :x:          |
| Markdown output    |         :x:          | :heavy_check_mark: |         :x:         |         :x:          |        :x:         |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          |         :x:          |         :x:          |

## Benchmark
The benchmark is taken from [simpletable] and adjusted for the libraries, if necessary.

All libraries show similar characteristics in terms of number of rows i.e., the time increases linearly with the number of rows (tested up to 1000).
However, the performance might be different for dozens of columns.
Furthermore, the impact of advanced features such as cell spanning, multi-line, custom formatters, etc. are not reflected in the benchmark since those features are not implemented by all libraries.
Thus, the benchmark implements a rather "simple" use case of five columns with strings and numbers.

The benchmark can be executed as follows:

```bash
go test ./... -bench . -benchmem -benchtime 10000x
```

```
?       github.com/gschauer/go-table-comparison [no test files]
goos: linux
goarch: amd64
pkg: github.com/gschauer/go-table-comparison/simple
BenchmarkGoPretty-8                10000            102503 ns/op           11940 B/op        312 allocs/op
BenchmarkGoPrettyTable-8           10000             54670 ns/op           11808 B/op        151 allocs/op
BenchmarkGoTabulate-8              10000             55644 ns/op           38108 B/op        546 allocs/op
BenchmarkGoTextTable-8             10000             98307 ns/op           24851 B/op        669 allocs/op
BenchmarkSimpleTable-8             10000            188199 ns/op           39531 B/op       1087 allocs/op
BenchmarkTabby-8                   10000             16833 ns/op            9139 B/op         74 allocs/op
BenchmarkTable-8                   10000             24369 ns/op            8275 B/op        244 allocs/op
BenchmarkTablePrinter-8            10000             18854 ns/op           11075 B/op         93 allocs/op
BenchmarkTableWriter-8             10000            140324 ns/op           22250 B/op        821 allocs/op
BenchmarkTabWriter-8               10000             14941 ns/op            8098 B/op         40 allocs/op
BenchmarkUITable-8                 10000            371959 ns/op           45158 B/op       2405 allocs/op
PASS
ok      github.com/gschauer/go-table-comparison/simple   11.064s
```

## Example Output
The output of many libraries can be highly customized.
The following listing shows whether they have sensible defaults in terms of output configuration.

In order to see the default output format of all implementations, run

```bash
go run main.go
```

### text/tabwriter
```
#     NAME              PHONE         EMAIL                        QTTY
----  ----              ----          ----                         ----
1     Newton G. Goetz   252-585-5166  NewtonGGoetz@dayrep.com      10
2     Rebecca R. Edney  865-475-4171  RebeccaREdney@armyspy.com    12
3     John R. Jackson   810-325-1417  JohnRJackson@armyspy.com     15
```

### github.com/jedib0t/go-pretty
```
+----+------------------+--------------+-----------------------------+------+
|  # | NAME             | PHONE        | EMAIL                       | QTTY |
+----+------------------+--------------+-----------------------------+------+
|  1 | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     |   10 |
|  2 | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   |   12 |
|  3 | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    |   15 |
+----+------------------+--------------+-----------------------------+------+
```

### github.com/tatsushid/go-prettytable
```
#  | NAME             | PHONE        | EMAIL                       | QTTY
1  | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     | 10
2  | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   | 12
3  | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    | 15
```

### github.com/bndr/gotabulate
```
+-------+---------------------+-----------------+--------------------------------+---------+
|     # |                NAME |           PHONE |                          EMAIL |    QTTY |
+=======+=====================+=================+================================+=========+
|     1 |     Newton G. Goetz |    252-585-5166 |        NewtonGGoetz@dayrep.com |      10 |
+-------+---------------------+-----------------+--------------------------------+---------+
|     2 |    Rebecca R. Edney |    865-475-4171 |      RebeccaREdney@armyspy.com |      12 |
+-------+---------------------+-----------------+--------------------------------+---------+
|     3 |     John R. Jackson |    810-325-1417 |       JohnRJackson@armyspy.com |      15 |
+-------+---------------------+-----------------+--------------------------------+---------+
```

### github.com/syohex/go-texttable
```
+----+------------------+--------------+-----------------------------+------+
| #  | NAME             | PHONE        | EMAIL                       | QTTY |
+----+------------------+--------------+-----------------------------+------+
|  1 | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     |   10 |
|  2 | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   |   12 |
|  3 | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    |   15 |
+----+------------------+--------------+-----------------------------+------+
```

### github.com/alexeyco/simpletable
```
+----+------------------+--------------+-----------------------------+------+
| #  | NAME             | PHONE        | EMAIL                       | QTTY |
+----+------------------+--------------+-----------------------------+------+
| 1  | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     | 10   |
| 2  | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   | 12   |
| 3  | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    | 15   |
+----+------------------+--------------+-----------------------------+------+
```

### github.com/cheynewallace/tabby
```
#   NAME              PHONE         EMAIL                        QTTY
-   ----              -----         -----                        ----
1   Newton G. Goetz   252-585-5166  NewtonGGoetz@dayrep.com      10
2   Rebecca R. Edney  865-475-4171  RebeccaREdney@armyspy.com    12
3   John R. Jackson   810-325-1417  JohnRJackson@armyspy.com     15
```

### github.com/rodaine/table
```
#   NAME              PHONE         EMAIL                        QTTY
1   Newton G. Goetz   252-585-5166  NewtonGGoetz@dayrep.com      10
2   Rebecca R. Edney  865-475-4171  RebeccaREdney@armyspy.com    12
3   John R. Jackson   810-325-1417  JohnRJackson@armyspy.com     15
```

### github.com/palantir/pkg/tableprinter
```
#   NAME              PHONE         EMAIL                        QTTY
-   ----              -----         -----                        ----
1   Newton G. Goetz   252-585-5166  NewtonGGoetz@dayrep.com      10
2   Rebecca R. Edney  865-475-4171  RebeccaREdney@armyspy.com    12
3   John R. Jackson   810-325-1417  JohnRJackson@armyspy.com     15
```

### github.com/olekukonko/tablewriter
```
+----+------------------+--------------+-----------------------------+------+
| #  |       NAME       |    PHONE     |            EMAIL            | QTTY |
+----+------------------+--------------+-----------------------------+------+
|  1 | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     |   10 |
|  2 | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   |   12 |
|  3 | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    |   15 |
+----+------------------+--------------+-----------------------------+------+
```

### github.com/gosuri/uitable
```
#       NAME                    PHONE           EMAIL                           QTTY
1       Newton G. Goetz         252-585-5166    NewtonGGoetz@dayrep.com         10
2       Rebecca R. Edney        865-475-4171    RebeccaREdney@armyspy.com       12
3       John R. Jackson         810-325-1417    JohnRJackson@armyspy.com        15
```

## Disclaimer
In case something is missing or incorrect, please file an issue or raise a pull request.

[tabwriter]: https://golang.org/pkg/text/tabwriter
[go-pretty]: https://github.com/jedib0t/go-pretty/tree/master/table
[go-prettytable]: https://github.com/tatsushid/go-prettytable
[gotabulate]: https://github.com/bndr/gotabulate
[go-texttable]: https://github.com/syohex/go-texttable
[simpletable]: https://github.com/alexeyco/simpletable
[tabby]: https://github.com/cheynewallace/tabby
[table]: https://github.com/rodaine/table
[tableprinter]: https://github.com/palantir/pkg
[tablewriter]: https://github.com/olekukonko/tablewriter
[uitable]: https://github.com/gosuri/uitable

[s_tabwriter]: https://img.shields.io/github/stars/golang/go?style=plastic&cacheSeconds=3600
[s_go-pretty]: https://img.shields.io/github/stars/jedib0t/go-pretty?style=plastic&cacheSeconds=3600
[s_go-prettytable]: https://img.shields.io/github/stars/tatsushid/go-prettytable?style=plastic&cacheSeconds=3600
[s_gotabulate]: https://img.shields.io/github/stars/bndr/gotabulate?style=plastic&cacheSeconds=3600
[s_go-texttable]: https://img.shields.io/github/stars/syohex/go-texttable?style=plastic&cacheSeconds=3600
[s_simpletable]: https://img.shields.io/github/stars/alexeyco/simpletable?style=plastic&cacheSeconds=3600
[s_tabby]: https://img.shields.io/github/stars/cheynewallace/tabby?style=plastic&cacheSeconds=3600
[s_table]: https://img.shields.io/github/stars/rodaine/table?style=plastic&cacheSeconds=3600
[s_tableprinter]: https://img.shields.io/github/stars/palantir/pkg?style=plastic&cacheSeconds=3600
[s_tablewriter]: https://img.shields.io/github/stars/olekukonko/tablewriter?style=plastic&cacheSeconds=3600
[s_uitable]: https://img.shields.io/github/stars/gosuri/uitable?style=plastic&cacheSeconds=3600

[c_tabwriter]: https://img.shields.io/github/last-commit/golang/go?style=plastic&cacheSeconds=3600
[c_go-pretty]: https://img.shields.io/github/last-commit/jedib0t/go-pretty?style=plastic&cacheSeconds=3600
[c_go-prettytable]: https://img.shields.io/github/last-commit/tatsushid/go-prettytable?style=plastic&cacheSeconds=3600
[c_gotabulate]: https://img.shields.io/github/last-commit/bndr/gotabulate?style=plastic&cacheSeconds=3600
[c_go-texttable]: https://img.shields.io/github/last-commit/syohex/go-texttable?style=plastic&cacheSeconds=3600
[c_simpletable]: https://img.shields.io/github/last-commit/alexeyco/simpletable?style=plastic&cacheSeconds=3600
[c_tabby]: https://img.shields.io/github/last-commit/cheynewallace/tabby?style=plastic&cacheSeconds=3600
[c_table]: https://img.shields.io/github/last-commit/rodaine/table?style=plastic&cacheSeconds=3600
[c_tableprinter]: https://img.shields.io/github/last-commit/palantir/pkg?style=plastic&cacheSeconds=3600
[c_tablewriter]: https://img.shields.io/github/last-commit/olekukonko/tablewriter?style=plastic&cacheSeconds=3600
[c_uitable]: https://img.shields.io/github/last-commit/gosuri/uitable?style=plastic&cacheSeconds=3600
