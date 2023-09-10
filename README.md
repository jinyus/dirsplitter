# dirsplitter

Split large directories into parts of a specified maximum size.

## How to run:

Download the binary from the [releases page](https://github.com/jinyus/dirsplitter/releases)

```bash
dirsplitter --help
```

## USAGE:

```text
Usage: dirsplitter [OPTIONS] COMMAND [arg...]

Split Directories

Options:
  -v, --version   Show the version and exit

Commands:
  split           Split a directory into parts of a given size
  reverse         Reverse a splitted directory

Run 'dirsplitter COMMAND --help' for more information on a command.
```

## SPLIT USAGE:

```text
Usage: dirsplitter split [OPTIONS] DIRECTORY

Split a directory into parts of a given size

Arguments:
  DIRECTORY    the directory to split (default ".")

Options:
  -m, --max    Size of each part in GB (default 5)
```

### example:

```bash
dirsplitter split --max 0.5 ./mylarge2GBdirectory

This will yield the following directory structure:

📂mylarge2GBdirectory
 |- 📂part1
 |- 📂part2
 |- 📂part3
 |- 📂part4

with each part being a maximum of 500MB in size.
```

Undo splitting

```bash
dirsplitter reverse ./mylarge2GBdirectory

```
