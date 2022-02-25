/**
 * Filename: fqcheck.go
 * Path: code/go
 * Created Date: Friday, February 25th 2022, 3:23:45 pm
 * Author: yibin
 *
 * Copyright (c) 2022 Yibin Wang
 */
package main

import (
	"io"
	"os"
	"github.com/op/go-logging"
	"github.com/shenwei356/bio/seqio/fastx"
	"github.com/shenwei356/xopen"
)

var log = logging.MustGetLogger("main")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05} %{shortfunc} | %{level:.6s} %{color:reset} %{message}`,
)

func main() {
	args := os.Args[1:]
	if len(args) < 2{
		print("Usage: fqcheck <input.fq.gz> <output.fq.gz>\n")
		print("    remove reads with larger than 100k\n")
		print("	and different length between seq and qual\n")
		os.Exit(3)
	}
	inputFastq := args[0]
	outFile := args[1]
	// import fastq
	log.Noticef("Loading fastq `%s`", inputFastq)
	reader, err := fastx.NewDefaultReader(inputFastq)
	if err != nil {
		log.Errorf("", err)
	}
	outfh, err := xopen.Wopen(outFile)
	defer outfh.Close()

	for {
		rec, err := reader.Read()
		if err == io.EOF || rec == nil{
			break
		}
		if rec.Seq.Length() != len(rec.Seq.Qual){
			continue
		}
		if rec.Seq.Length() > 100000{
			continue
		}
		rec.FormatToWriter(outfh, 0)
	}
	log.Noticef("Done, output fastq `%s`", outFile)
}