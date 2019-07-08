# LazyPass

The CLI tool that generates random password for lazy dev. Inspired by [radomix by kkdai](https://github.com/kkdai/radomtix)

## Install

```
go get github.com/tatthien/lazypass
```

## Usage

```
$ lazypass

KCncmRwhEw
```

```
$ lazypass -n -s -l 20 

1B^Ct0H5br^EQ9=c1SRf
```

```
$ lazypass -n -l 20

SC5DZssBJINpoTKtK0Q4
```

```
$ lazypass -s -l 20 

=OSC%u^$h!Pt&vL_]rJ!
```

## Options

- `-l` length of password (default is 10)
- `-n` password contains numbers (default false)
- `-s` password contains symbols (default false)