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

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
