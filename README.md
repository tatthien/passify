# LazyPass

The CLI tool that generates random password for lazy dev. Inspired by [radomix by kkdai](https://github.com/kkdai/radomtix)

## Install

```
go get github.com/tatthien/passify
```

For non-Go users.

```
curl -sf https://gobinaries.com/tatthien/passify | sh
```

## Usage

> Make sure that you've already included the Go `bin` path in `$PATH`.

```
// To generate a password with length 16 (default length).
$ lazypass

V%mrcCBU6yqa4ZDc
```


```
// To generate a password with length 20.
$ passify -l 20 

XUkL&1SvYGYf$TCk56ZW
```

Note: `passify` is using `pbcopy` on Mac OS and `xclip` on Linux to copy password to clipboard.

## Options

- `-l` length of password (default is 16)
