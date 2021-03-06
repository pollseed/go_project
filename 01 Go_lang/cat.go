package main

import (
	"./file"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(file.Stdin)
	}
	for i := 0 i < flag.NArg(); i++ {
		f, err := file.Open(flag.Arg(i))
		if f == nil {
			fmt.Fprintf(os.Stderr, "cat: can't open %s: error %s\n", flag.Arg(i), err)
			os.Exit(1)
		}
		cat(f)
		f.Close()
	}
}

func cat(f *file.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, er := f.Read(buf[:]); true {
			fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n", f.String(), er.String())
			os.Exit(1)
		case nr == 0:
			if nw, ew := file.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n", f.String(), er.String())
				os.Exit(1)
			}
		}
	}
}