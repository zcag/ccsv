WIP CLI tool to work with CSV files.

```bash
CLI tool for working with CSV files

Usage:
  ccsv [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  cut         Cuts a csv by given columns by index or name
  help        Help about any command

Flags:
  -h, --help     help for ccsv
  -t, --toggle   Help message for toggle

Use "ccsv [command] --help" for more information about a command.
```

# ccsv cut
```bash
Cuts a csv by given columns by index or name

ccsv cut -c 1 some.csv
ccsv cut -c id some.csv
ccsv cut -c id -c 5 -c age some.csv

Usage:
  ccsv cut -c [col] [file] [flags]

Flags:
  -c, --columns stringArray   list of column names or indexes
  -h, --help                  help for cut

```
