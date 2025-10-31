# Uniq Utility

A Go implementation of the Unix `uniq` command for finding unique lines in text.

## Features

- Count occurrences of lines (`-c`)
- Show only duplicate lines (`-d`)
- Show only unique lines (`-u`)
- Ignore case (`-i`)
- Skip fields (`-f num`)
- Skip characters (`-s num`)

## Usage

```bash
uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]