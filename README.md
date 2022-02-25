# fqcheck: Check Pore-C reads

## Introduction 

`fqcheck` can be used to filter the Pore-C reads which are easily get error of reads length larger 100k or with different length of seq and qual.  

## Installation
The easiest way to install fqcheck is to download the latest binary from the [releases](https://github.com/wangyibin/fqcheck/releases/latest)  

Build from source with:
```bash
go get -u -t -v github.com/wangyibin/fqcheck/...
go install github.com/wangyibin/fqcheck/fqcheck
```

## Usage
```bash
fqcheck input.fq.gz output.fq.gz
```
