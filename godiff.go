package godiff

import (
	"math"
)

type Diff struct {
	str1 string
	str2 string
}

func NewDiff(s1, s2 string) (*Diff, error) {
	// if len(s1) >= len(s2) { // set s2 as longer string
	// 	s1, s2 = s2, s1
	// }

	return &Diff{s1, s2}, nil
}

func (d *Diff) LevenshteinDistance() int {
	l1 := len(d.str1)
	l2 := len(d.str2)

	dist := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dist[i] = make([]int, l2+1)

		dist[i][0] = i
	}
	for j := 0; j <= l2; j++ {
		dist[0][j] = j
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if d.str1[i-1] == d.str2[j-1] {
				dist[i][j] = min(dist[i-1][j-1], min(dist[i][j-1]+1, dist[i-1][j]+1))
			} else {
				dist[i][j] = min(dist[i][j-1]+1, dist[i-1][j]+1)
			}
		}
	}

	return dist[l1][l2]
}

func (d *Diff) LongCommonSubSeq() string {
	// compute lcs table
	l := computeLCSTable(d.str1, d.str2)

	// debug
	// for i := 0; i <= len(d.str1); i++ {
	// 	for j := 0; j < len(d.str2); j++ {
	// 		print(l[i][j])
	// 		print(" ")
	// 	}
	// 	print(l[i][len(d.str2)])
	// 	println()
	// }

	// assemble lcs
	return assembleLCS(d.str1, d.str2, l, len(d.str1), len(d.str2))
}

func computeLCSTable(s1, s2 string) [][]int {
	l := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		l[i] = make([]int, len(s2)+1)
		l[i][0] = 0
	}
	for j := 0; j <= len(s2); j++ {
		l[0][j] = 0
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				l[i][j] = l[i-1][j-1] + 1
			} else {
				l[i][j] = max(l[i][j-1], l[i-1][j])
			}
		}
	}

	return l
}

func assembleLCS(s1, s2 string, l [][]int, i, j int) string {
	if l[i][j] == 0 {
		return ""
	}

	if s1[i-1] == s2[j-1] {
		return assembleLCS(s1, s2, l, i-1, j-1) + string(s1[i-1])
	} else {
		if l[i][j-1] > l[i-1][j] {
			return assembleLCS(s1, s2, l, i, j-1)
		} else {
			return assembleLCS(s1, s2, l, i-1, j)
		}
	}
}

type Op int

const (
	cop Op = iota // copy
	rep           // replace
	ins           // insert
	del           // delete
)

func (d *Diff) Transform() []string {
	// compute transformation table
	_, op := computeTransformTable(d.str1, d.str2) // cost for debug

	// debug
	// for i := 0; i <= len(d.str1); i++ {
	// 	for j := 0; j < len(d.str2); j++ {
	// 		print(cost[i][j])
	// 		print(" ")
	// 	}
	// 	print(cost[i][len(d.str2)])
	// 	println()
	// }

	ops := make([]string, 0)

	// assemble transformation
	return assembleTransform(op, d.str1, d.str2, len(d.str1), len(d.str2), ops)
}

func computeTransformTable(s1, s2 string) ([][]int, [][]Op) {
	cost := make([][]int, len(s1)+1)
	op := make([][]Op, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		cost[i] = make([]int, len(s2)+1)
		op[i] = make([]Op, len(s2)+1)

		cost[i][0] = i * 2
		op[i][0] = del
	}
	for j := 0; j <= len(s2); j++ {
		cost[0][j] = j * 2
		op[0][j] = ins
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			// check copy or replace
			if s1[i-1] == s2[j-1] {
				cost[i][j] = cost[i-1][j-1] - 1
				op[i][j] = cop
			} else {
				cost[i][j] = cost[i-1][j-1] + 1
				op[i][j] = rep
			}

			if cost[i-1][j]+2 < cost[i][j] {
				cost[i][j] = cost[i-1][j] + 2
				op[i][j] = del
			}
			if cost[i][j-1]+2 < cost[i][j] {
				cost[i][j] = cost[i][j-1] + 2
				op[i][j] = ins
			}
		}
	}

	return cost, op
}

func assembleTransform(op [][]Op, s1, s2 string, i, j int, ops []string) []string {
	if i == 0 && j == 0 {
		return nil
	}

	if op[i][j] == cop {
		return append(assembleTransform(op, s1, s2, i-1, j-1, ops), "copy "+string(s1[i-1]))
	} else if op[i][j] == rep {
		return append(assembleTransform(op, s1, s2, i-1, j-1, ops), "replace "+string(s1[i-1])+" by "+string(s2[j-1]))
	} else {
		if op[i][j] == del {
			return append(assembleTransform(op, s1, s2, i-1, j, ops), "delete "+string(s1[i-1]))
		} else { // op[i][j] == ins
			return append(assembleTransform(op, s1, s2, i, j-1, ops), "insert "+string(s2[j-1]))
		}
	}
}

// ShowDiff .
// show like below
//
//  str1: ATGATCG_GCAT_
//  str2: _CAAT_GTGAATC
//  diff: *++--+-+-+--+
func (d *Diff) ShowDiff() []string {
	// compute transformation table
	_, op := computeTransformTable(d.str1, d.str2) // cost for debug

	// assembleDiff
	var dstr1 string
	var dstr2 string
	var diffs string
	dstr1 = assembleDiffStr(op, d.str1, d.str2, len(d.str1), len(d.str2), dstr1, true)
	dstr2 = assembleDiffStr(op, d.str1, d.str2, len(d.str1), len(d.str2), dstr2, false)
	diffs = assembleDiff(op, d.str1, d.str2, len(d.str1), len(d.str2), diffs)

	return []string{dstr1, dstr2, diffs}
}

func assembleDiffStr(op [][]Op, s1, s2 string, i, j int, ret string, before bool) string {
	if i == 0 && j == 0 {
		return ""
	}

	if op[i][j] == cop {
		return assembleDiffStr(op, s1, s2, i-1, j-1, ret, before) + string(s1[i-1])
	} else if op[i][j] == rep {
		if before {
			return assembleDiffStr(op, s1, s2, i-1, j-1, ret, before) + string(s1[i-1])
		} else {
			return assembleDiffStr(op, s1, s2, i-1, j-1, ret, before) + string(s2[j-1])
		}
	} else {
		if op[i][j] == del {
			if before {
				return assembleDiffStr(op, s1, s2, i-1, j, ret, before) + string(s1[i-1])
			} else {
				return assembleDiffStr(op, s1, s2, i-1, j, ret, before) + "_"
			}
		} else { // op[i][j] == ins
			if before {
				return assembleDiffStr(op, s1, s2, i, j-1, ret, before) + "_"
			} else {
				return assembleDiffStr(op, s1, s2, i, j-1, ret, before) + string(s2[j-1])
			}
		}
	}
}

func assembleDiff(op [][]Op, s1, s2 string, i, j int, ret string) string {
	if i == 0 && j == 0 {
		return ""
	}

	if op[i][j] == cop {
		return assembleDiff(op, s1, s2, i-1, j-1, ret) + "-"
	} else if op[i][j] == rep {
		return assembleDiff(op, s1, s2, i-1, j-1, ret) + "+"
	} else {
		if op[i][j] == del {
			return assembleDiff(op, s1, s2, i-1, j, ret) + "*"
		} else { // op[i][j] == ins
			return assembleDiff(op, s1, s2, i, j-1, ret) + "*"
		}
	}
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
