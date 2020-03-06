# LazyPass

The CLI tool that generates random password for lazy dev. Inspired by [radomix by kkdai](https://github.com/kkdai/radomtix)

## Install

```
go get github.com/tatthien/passify
```

## Usage

> Make sure that you've already included the Go `bin` path in `$PATH`.

```
// To generate a password with length 16 (default length).
$ lazypass

E!I$^3@v^QlCwacD
```


```
// To generate a password with length 20.
$ passify -l 20 

HI7$A5-4A[xwIx?sns6Q
```

```
// To generate a password and copy to clipboard.
$ passify -c

Copied to clipboard
```

Note: `passify` is using `pbcopy` on Mac OS and `xclip` on Linux to copy password to clipboard.

## Options

- `-l` length of password (default is 10)
- `-c` copy password to clipboard (default is false)
