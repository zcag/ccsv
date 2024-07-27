# CCSV

CLI tool for working with CSV files. 
Personal project for getting into golang.

Only unique feature is to ability to diff csv files based on arbitrary columns.

For the rest you're probably better off with [csvkit](https://csvkit.readthedocs.io/en/latest/)

## Installation
```bash
go install github.com/cagdassalur/ccsv@latest
```

## Usage

Run `ccsv --help` or `ccsv [commmand] --help` for complete usage. 

* All column flags can be either column index or header name
* All commands except diff can work with piped csv input or from a positional file argument.

### diff
Provides a csv with unique rows from left side, filtered by an arbitrary column.

```bash
ccsv diff -l 1 -r 4 left.csv right.csv
ccsv diff -l id -r user_id left.csv right.csv
ccsv diff -c id left.csv right.csv
```

### headers
List headers and their indexes of file

```bash
ccsv headers some.csv
```

### cut
Filter csv to show specified columns

```bash
ccsv cut -c 1 some.csv
ccsv cut -c id some.csv
cat some.csv | ccsv cut -c id -c 5 -c age
```

### stat
Show info about each column like data type, null count, uniq count, min/max/mean/summ

```bash
ccsv stat some.csv
ccsv stat -H headerless.csv
```
