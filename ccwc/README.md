# CCWC (Code Challenges Word Counter)

## Description

CCWC is a CLI tool developed in Go, designed to perform basic text file analyses. It can count lines, words, characters, and bytes in a given text file.

## Features

- Count the number of lines in a file.
- Count the number of words in a file.
- Count the number of characters in a file.
- Count the number of bytes in a file.

## Installation

To install CCWC, clone the repository and build the project using Go.

```bash
git clone <repository-url>
cd ccwc
go build
```

## Usage

Run the tool using the command line. Here are some examples:

```bash
./ccwc -l <fileName>   # Count lines
./ccwc -w <fileName>   # Count words
./ccwc -c <fileName>   # Count bytes
./ccwc -m <fileName>   # Count characters
```
