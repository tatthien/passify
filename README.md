# LazyPass

The CLI tool that generates random password for lazy dev. Inspired by [radomix by kkdai](https://github.com/kkdai/radomtix)

## Install

```
go get github.com/tatthien/lazypass
```

## Usage

```
// To generate password with length 10 that contains letters only.
$ lazypass

KCncmRwhEw
```

```
// To generate password with length 20 that contains letters, numbers and symbols.
$ lazypass -n -s -l 20 

1B^Ct0H5br^EQ9=c1SRf
```

```
// To generate password with length 20 that contains letters and numbers.
$ lazypass -n -l 20

SC5DZssBJINpoTKtK0Q4
```

```
// To generate password with length 20 that contains letters and symbols.
$ lazypass -s -l 20 

=OSC%u^$h!Pt&vL_]rJ!
```

```
// To generate password and copy to clipboard.
$ lazypass -c

Copied to clipboard
```

Note: `lazypass` is using `pbcopy` on Mac OS and `xclip` on Linux to copy password to clipboard.

## Options

- `-l` length of password (default is 10)
- `-n` password contains numbers (default false)
- `-s` password contains symbols (default false)
- `-c` copy password to clipboard (default false)