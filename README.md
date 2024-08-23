# CCSV

WIP CLI tool for working with csv files

Personal project for getting into golang.

## Features
* Produce diff.csv based specific columns from 2 files
* Cut some columns to output a smaller width csv file
* Shows stats on csv file like data type, null count, min/max/mean/sum
* Filter rows by applying regex pattern to given columns
* Produce multiple csv files grouped by given column values
* header-skip, header-restore for hassle free csv piping

## Installation
```bash
go install github.com/zcag/ccsv@latest
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

### header-skip header-restore
Strip headers and restore them later to easily work on csv files without touching header line

```bash
cat file.csv | ccsv header-skip | sort | ccsv header-restore
cat file.csv | ccsv hs | sort | ccsv hr
```

### match
Filter rows by regex on specific columns

```bash
ccsv match -c name '\w+_\d' some.csv
cat some.csv | ccsv match -c email '.*@g?mail'
```

### group
Create csv files grouped by specified column while preserving headers

ex. from `customers.csv` you can get `TR_customers.csv, EU_customers.csv etc.`

```bash
ccsv group 'records_<country>' all_records.csv
cat some.csv | ccsv group '<3>_records_grouped'`,
```
