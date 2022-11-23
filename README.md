# Go Table Library Comparison

There are several libraries for rending tables in console applications written Go.
However, it's difficult to compare their features as well as their performance.

Below you can find an subjective overview of several table rendering libraries:

* [golang.org/pkg/text/tabwriter][tabwriter]
* [github.com/alexeyco/simpletable][alexeyco]
* [github.com/bndr/gotabulate][bndr]
* [github.com/cheynewallace/tabby][cheynewallace]
* [github.com/gosuri/uitable][gosuri]
* [github.com/jedib0t/go-pretty][jedib0t]
* [github.com/olekukonko/tablewriter][olekukonko]
* [github.com/palantir/pkg][palantir]
* [github.com/rodaine/table][rodaine]
* [github.com/syohex/go-texttable][syohex]
* [github.com/tatsushid/go-prettytable][tatsushid]

## Feature Comparison

Please note that the absence of a feature does not mean that there is no support at all.

For example, many libraries don't have color support built-in but are compatible with common ANSI color libraries.

| Feature            |     [tabwriter]      |      [alexeyco]      |        [bndr]        |      [caarlos0]      |   [cheynewallace]    |       [gosuri]       |     [jedib0t]      |     [olekukonko]     |      [palantir]      |      [rodaine]       |      [syohex]      |     [tatsushid]     |
| :----------------- | :------------------: | :------------------: | :------------------: | :------------------: | :------------------: | :------------------: | :----------------: | :------------------: | :------------------: | :------------------: | :----------------: | :-----------------: |
| Github last commit |                      |    ![c_alexeyco]     |      ![c_bndr]       |    ![c_caarlos0]     |  ![c_cheynewallace]  |     ![c_gosuri]      |    ![c_jedib0t]    |   ![c_olekukonko]    |    ![c_palantir]     |     ![c_rodaine]     |    ![c_syohex]     |   ![c_tatsushid]    |
| Github popularity  |  language built-in   |    ![s_alexeyco]     |      ![s_bndr]       |    ![s_caarlos0]     |  ![s_cheynewallace]  |     ![s_gosuri]      |    ![s_jedib0t]    |   ![s_olekukonko]    |    ![s_palantir]     |     ![s_rodaine]     |    ![s_syohex]     |   ![s_tatsushid]    |
|                    |                      |                      |                      |                      |                      |                      |                    |                      |                      |                      |                    |                     |
| performance        |  :heavy_check_mark:  | :small_blue_diamond: |  :heavy_check_mark:  | :small_blue_diamond: |  :heavy_check_mark:  |         :x:          | :heavy_check_mark: | :small_blue_diamond: |  :heavy_check_mark:  |  :heavy_check_mark:  | :heavy_check_mark: | :heavy_check_mark:  |
|                    |                      |                      |                      |                      |                      |                      |                    |                      |                      |                      |                    |                     |
| built-in colors    |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          |         :x:          |         :x:          | :heavy_check_mark: |  :heavy_check_mark:  |         :x:          |         :x:          |        :x:         |         :x:         |
| predefined styles  |         :x:          |  :heavy_check_mark:  |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          | :heavy_check_mark: |         :x:          |         :x:          |         :x:          |        :x:         |         :x:         |
| custom styles      |         :x:          |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          |         :x:          | :heavy_check_mark: |  :heavy_check_mark:  |         :x:          |         :x:          |        :x:         | :heavy_check_mark:  |
|                    |                      |                      |                      |                      |                      |                      |                    |                      |                      |                      |                    |                     |
| align text         | :small_blue_diamond: |  :heavy_check_mark:  | :small_blue_diamond: | :small_blue_diamond: | :small_blue_diamond: | :small_blue_diamond: | :heavy_check_mark: |  :heavy_check_mark:  | :small_blue_diamond: |         :x:          |        :x:         | :heavy_check_mark:  |
| borders            | :small_blue_diamond: |  :heavy_check_mark:  |  :heavy_check_mark:  | :small_blue_diamond: | :small_blue_diamond: |         :x:          | :heavy_check_mark: |  :heavy_check_mark:  | :small_blue_diamond: |         :x:          |        :x:         |         :x:         |
| cell span          |         :x:          |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          |         :x:          | :heavy_check_mark: |  :heavy_check_mark:  |         :x:          |         :x:          |        :x:         |         :x:         |
| custom renderer    |         :x:          |  :heavy_check_mark:  |         :x:          |  :heavy_check_mark:  |         :x:          |         :x:          | :heavy_check_mark: |  :heavy_check_mark:  |         :x:          | :small_blue_diamond: |        :x:         |         :x:         |
| footer             |         :x:          |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          |         :x:          | :heavy_check_mark: |  :heavy_check_mark:  |         :x:          |         :x:          |        :x:         |         :x:         |
| multiline          |         :x:          |  :heavy_check_mark:  |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          | :heavy_check_mark: |  :heavy_check_mark:  |         :x:          |         :x:          |        :x:         |         :x:         |
|                    |                      |                      |                      |                      |                      |                      |                    |                      |                      |                      |                    |                     |
| string values      |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  | :heavy_check_mark: |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  | :heavy_check_mark: | :heavy_check_mark:  |
| interface values   |  :heavy_check_mark:  |         :x:          |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  | :heavy_check_mark: |         :x:          |  :heavy_check_mark:  |  :heavy_check_mark:  |        :x:         | :heavy_check_mark:  |
| arbitrary structs  |         :x:          |         :x:          |         :x:          |  :heavy_check_mark:  |         :x:          |         :x:          |        :x:         |         :x:          |  :heavy_check_mark:  |         :x:          |        :x:         |         :x:         |
|                    |                      |                      |                      |                      |                      |                      |                    |                      |                      |                      |                    |                     |
| ASCII output       |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  | :heavy_check_mark: |  :heavy_check_mark:  |  :heavy_check_mark:  |  :heavy_check_mark:  | :heavy_check_mark: | :heavy_check_mark:  |
| CSV output         |         :x:          |         :x:          |         :x:          |         :x:          |         :x:          |         :x:          | :heavy_check_mark: |         :x:          |         :x:          |         :x:          |        :x:         |         :x:         |
| HTML output        |         :x:          |         :x:          |         :x:          |         :x:          |         :x:          |         :x:          | :heavy_check_mark: |         :x:          |         :x:          |         :x:          |        :x:         |         :x:         |
| Markdown output    |         :x:          |  :heavy_check_mark:  |         :x:          |         :x:          |         :x:          |         :x:          | :heavy_check_mark: |         :x:          |         :x:          |         :x:          |        :x:         |         :x:         |

## Benchmark

The benchmark is taken from [alexeyco] and adjusted for the libraries, if necessary.

All libraries show similar characteristics in terms of number of rows i.e., the time increases linearly with the number of rows (tested up to 1000).
However, the performance might be different for dozens of columns.
Furthermore, the impact of advanced features such as cell spanning, multi-line, custom formatters, etc. are not reflected in the benchmark since those features are not implemented by all libraries.
Thus, the benchmark implements a rather "simple" use case of five columns with strings and numbers.

The benchmark can be executed as follows:

```bash
go test ./... -bench . -benchmem -benchtime 10000x
```

```text
go test ./... -bench . -benchmem -benchtime 10000x
?       github.com/gschauer/go-table-comparison [no test files]
goos: darwin
goarch: arm64
pkg: github.com/gschauer/go-table-comparison/simple
BenchmarkAlexeyCo-10               10000            125782 ns/op           63461 B/op       1418 allocs/op
BenchmarkBndr-10                   10000             23313 ns/op           36994 B/op        536 allocs/op
BenchmarkCaarlos0-10               10000              6747 ns/op            8708 B/op         86 allocs/op
BenchmarkCheyneWallace-10          10000              7392 ns/op            8708 B/op         72 allocs/op
BenchmarkGOsuri-10                 10000            276073 ns/op          117436 B/op       3532 allocs/op
BenchmarkJedib0t-10                10000             36386 ns/op           11005 B/op        310 allocs/op
BenchmarkOlekuKonko-10             10000            102845 ns/op           46892 B/op       1196 allocs/op
BenchmarkPalantir-10               10000              8057 ns/op           10677 B/op         93 allocs/op
BenchmarkRodaine-10                10000             10062 ns/op            8820 B/op        245 allocs/op
BenchmarkSyohex-10                 10000             82481 ns/op           47804 B/op        999 allocs/op
BenchmarkTabWriter-10              10000              5971 ns/op            7715 B/op         40 allocs/op
BenchmarkTatsushiD-10              10000             66908 ns/op           34240 B/op        481 allocs/op
PASS
ok      github.com/gschauer/go-table-comparison/simple  7.990s
```

## Example Output

The output of many libraries can be highly customized.
The following listing shows whether they have sensible defaults in terms of output configuration.

In order to see the default output format of all implementations, run

```bash
go run main.go
```

### text/tabwriter

```text
#     NAME              PHONE         EMAIL                        QTTY
----  ----              ----          ----                         ----
1     Newton G. Goetz   252-585-5166  NewtonGGoetz@dayrep.com      10
2     Rebecca R. Edney  865-475-4171  RebeccaREdney@armyspy.com    12
3     John R. Jackson   810-325-1417  JohnRJackson@armyspy.com     15
```

### github.com/alexeyco/simpletable

```text
+----+------------------+--------------+-----------------------------+------+
| #  | NAME             | PHONE        | EMAIL                       | QTTY |
+----+------------------+--------------+-----------------------------+------+
| 1  | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     | 10   |
| 2  | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   | 12   |
| 3  | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    | 15   |
+----+------------------+--------------+-----------------------------+------+
```

### github.com/bndr/gotabulate

```text
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

### github.com/caarlos0/tablewriter

```text
#        NAME                PHONE           EMAIL                          QTTY
1        Newton G. Goetz     252-585-5166    NewtonGGoetz@dayrep.com        10
2        Rebecca R. Edney    865-475-4171    RebeccaREdney@armyspy.com      12
3        John R. Jackson     810-325-1417    JohnRJackson@armyspy.com       15
```

### github.com/cheynewallace/tabby

```text
#   NAME              PHONE         EMAIL                        QTTY
-   ----              -----         -----                        ----
1   Newton G. Goetz   252-585-5166  NewtonGGoetz@dayrep.com      10
2   Rebecca R. Edney  865-475-4171  RebeccaREdney@armyspy.com    12
3   John R. Jackson   810-325-1417  JohnRJackson@armyspy.com     15
```

### github.com/gosuri/uitable

```text
#       NAME                    PHONE           EMAIL                           QTTY
1       Newton G. Goetz         252-585-5166    NewtonGGoetz@dayrep.com         10
2       Rebecca R. Edney        865-475-4171    RebeccaREdney@armyspy.com       12
3       John R. Jackson         810-325-1417    JohnRJackson@armyspy.com        15
```

### github.com/jedib0t/go-pretty

```text
+----+------------------+--------------+-----------------------------+------+
|  # | NAME             | PHONE        | EMAIL                       | QTTY |
+----+------------------+--------------+-----------------------------+------+
|  1 | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     |   10 |
|  2 | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   |   12 |
|  3 | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    |   15 |
+----+------------------+--------------+-----------------------------+------+
```

### github.com/olekukonko/tablewriter

```text
+----+------------------+--------------+-----------------------------+------+
| #  |       NAME       |    PHONE     |            EMAIL            | QTTY |
+----+------------------+--------------+-----------------------------+------+
|  1 | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     |   10 |
|  2 | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   |   12 |
|  3 | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    |   15 |
+----+------------------+--------------+-----------------------------+------+
```

### github.com/palantir/pkg/tableprinter

```text
#   NAME              PHONE         EMAIL                        QTTY
-   ----              -----         -----                        ----
1   Newton G. Goetz   252-585-5166  NewtonGGoetz@dayrep.com      10
2   Rebecca R. Edney  865-475-4171  RebeccaREdney@armyspy.com    12
3   John R. Jackson   810-325-1417  JohnRJackson@armyspy.com     15
```

### github.com/rodaine/table

```text
#   NAME              PHONE         EMAIL                        QTTY
1   Newton G. Goetz   252-585-5166  NewtonGGoetz@dayrep.com      10
2   Rebecca R. Edney  865-475-4171  RebeccaREdney@armyspy.com    12
3   John R. Jackson   810-325-1417  JohnRJackson@armyspy.com     15
```

### github.com/syohex/go-texttable

```text
+----+------------------+--------------+-----------------------------+------+
| #  | NAME             | PHONE        | EMAIL                       | QTTY |
+----+------------------+--------------+-----------------------------+------+
|  1 | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     |   10 |
|  2 | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   |   12 |
|  3 | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    |   15 |
+----+------------------+--------------+-----------------------------+------+
```

### github.com/tatsushid/go-prettytable

```text
#  | NAME             | PHONE        | EMAIL                       | QTTY
1  | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     | 10
2  | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   | 12
3  | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    | 15
```

## Disclaimer

In case something is missing or incorrect, please file an issue or raise a pull request.

[tabwriter]: https://golang.org/pkg/text/tabwriter
[alexeyco]: https://github.com/alexeyco/simpletable
[bndr]: https://github.com/bndr/gotabulate
[caarlos0]: https://github.com/caarlos0/tablewriter
[cheynewallace]: https://github.com/cheynewallace/tabby
[gosuri]: https://github.com/gosuri/uitable
[jedib0t]: https://github.com/jedib0t/go-pretty/tree/master/table
[olekukonko]: https://github.com/olekukonko/tablewriter
[palantir]: https://github.com/palantir/pkg
[rodaine]: https://github.com/rodaine/table
[syohex]: https://github.com/syohex/go-texttable
[tatsushid]: https://github.com/tatsushid/go-prettytable

[s_alexeyco]: https://img.shields.io/github/stars/alexeyco/simpletable?style=plastic&cacheSeconds=3600
[s_bndr]: https://img.shields.io/github/stars/bndr/gotabulate?style=plastic&cacheSeconds=3600
[s_caarlos0]: https://img.shields.io/github/stars/caarlos0/tablewriter?style=plastic&cacheSeconds=3600
[s_cheynewallace]: https://img.shields.io/github/stars/cheynewallace/tabby?style=plastic&cacheSeconds=3600
[s_gosuri]: https://img.shields.io/github/stars/gosuri/uitable?style=plastic&cacheSeconds=3600
[s_jedib0t]: https://img.shields.io/github/stars/jedib0t/go-pretty?style=plastic&cacheSeconds=3600
[s_olekukonko]: https://img.shields.io/github/stars/olekukonko/tablewriter?style=plastic&cacheSeconds=3600
[s_palantir]: https://img.shields.io/github/stars/palantir/pkg?style=plastic&cacheSeconds=3600
[s_rodaine]: https://img.shields.io/github/stars/rodaine/table?style=plastic&cacheSeconds=3600
[s_syohex]: https://img.shields.io/github/stars/syohex/go-texttable?style=plastic&cacheSeconds=3600
[s_tatsushid]: https://img.shields.io/github/stars/tatsushid/go-prettytable?style=plastic&cacheSeconds=3600

[c_alexeyco]: https://img.shields.io/github/last-commit/alexeyco/simpletable?style=plastic&cacheSeconds=3600
[c_bndr]: https://img.shields.io/github/last-commit/bndr/gotabulate?style=plastic&cacheSeconds=3600
[c_caarlos0]: https://img.shields.io/github/last-commit/caarlos0/tablewriter?style=plastic&cacheSeconds=3600
[c_cheynewallace]: https://img.shields.io/github/last-commit/cheynewallace/tabby?style=plastic&cacheSeconds=3600
[c_gosuri]: https://img.shields.io/github/last-commit/gosuri/uitable?style=plastic&cacheSeconds=3600
[c_jedib0t]: https://img.shields.io/github/last-commit/jedib0t/go-pretty?style=plastic&cacheSeconds=3600
[c_olekukonko]: https://img.shields.io/github/last-commit/olekukonko/tablewriter?style=plastic&cacheSeconds=3600
[c_palantir]: https://img.shields.io/github/last-commit/palantir/pkg?style=plastic&cacheSeconds=3600
[c_rodaine]: https://img.shields.io/github/last-commit/rodaine/table?style=plastic&cacheSeconds=3600
[c_syohex]: https://img.shields.io/github/last-commit/syohex/go-texttable?style=plastic&cacheSeconds=3600
[c_tatsushid]: https://img.shields.io/github/last-commit/tatsushid/go-prettytable?style=plastic&cacheSeconds=3600
