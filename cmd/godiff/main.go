package main

import (
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
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s str1 str2\n", name)
		return exitNG
	}

	diff, err := godiff.NewDiff(args[0], args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
		return exitNG
	}
	fmt.Fprintf(os.Stdout, "Distance: %d\n", diff.LevenshteinDistance())
	fmt.Fprintf(os.Stdout, "LCS: %s\n", diff.LongCommonSubSeq())
	fmt.Fprintf(os.Stdout, "Transformation %s to %s:\n", args[0], args[1])
	ops := diff.Transform()
	for _, op := range ops {
		fmt.Fprintf(os.Stdout, "\t%s\n", op)
	}

	return exitOK
}
