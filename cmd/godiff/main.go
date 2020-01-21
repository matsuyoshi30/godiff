package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/matsuyoshi30/godiff"
)

const name = "godiff"

func main() {
	os.Exit(run(os.Args[1:]))
}

const (
	exitOK = iota
	exitNG
)

func run(args []string) int {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	fs.Usage = func() {
		fmt.Fprintf(os.Stdout, "usage: %s str1 str2\n", name)
		flag.PrintDefaults()
	}

	var showAll bool
	var modeWord bool
	fs.BoolVar(&showAll, "all", false, "show all info")
	fs.BoolVar(&modeWord, "word", false, "word")
	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return exitOK
		}
		return exitNG
	}

	args = fs.Args()

	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s str1 str2\n", name)
		return exitNG
	}

	diff, err := godiff.NewDiff(args[0], args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
		return exitNG
	}

	if showAll {
		fmt.Fprintf(os.Stdout, "Distance: %d\n", diff.LevenshteinDistance())
		fmt.Fprintf(os.Stdout, "LCS: %s\n", diff.LongCommonSubSeq())
		fmt.Fprintf(os.Stdout, "Transformation %s to %s:\n", args[0], args[1])
		ops := diff.Transform()
		for _, op := range ops {
			fmt.Fprintf(os.Stdout, "\t%s\n", op)
		}
	} else if modeWord {
		diffs := diff.ShowDiff()
		fmt.Fprintf(os.Stdout, "BEFORE: %s\n", diffs[0])
		fmt.Fprintf(os.Stdout, "AFTER : %s\n", diffs[1])
		fmt.Fprintf(os.Stdout, "DIFF  : %s\n", diffs[2])
	} else { // file
		diffs, err := diff.ShowFileDiff()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
		}

		for _, d := range diffs {
			fmt.Fprintf(os.Stdout, "%s\n", d)
		}
	}

	return exitOK
}
