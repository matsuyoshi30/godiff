package godiff

import (
	"math"
)

type Diff struct {
	str1 string
	str2 string
}

func NewDiff(s1, s2 string) (*Diff, error) {
	if len(s1) >= len(s2) { // set s2 as longer string
		s1, s2 = s2, s1
	}

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

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
