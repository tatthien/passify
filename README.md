# LazyPass

The CLI tool that generates random password for lazy dev. Inspired by [radomix by kkdai](https://github.com/kkdai/radomtix)

## Install

```
go get github.com/tatthien/lazypass
```

## Usage

```
// To generate a password with length 16 (default length).
$ lazypass

KCncmRwhEw
```


```
// To generate a password with length 20.
$ lazypass -l 20 

=OSC%u^$h!Pt&vL_]rJ!
```

```
// To generate a password and copy to clipboard.
$ lazypass -c

Copied to clipboard
```

Note: `lazypass` is using `pbcopy` on Mac OS and `xclip` on Linux to copy password to clipboard.

## Options

- `-l` length of password (default is 10)
- `-c` copy password to clipboard (default is false)