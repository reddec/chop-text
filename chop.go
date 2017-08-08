package chop

import (
	"math"
)

// Chopper configuration
type Chopper struct {
	Slices    int  // Maximum count of slices in result string
	SliceSize int  // Size of each slice
	Sep       byte // Separator between slices
}

// Chop string to slices with specified slice count, size and separator
func (c *Chopper) Chop(text string) string {
	N := len(text)
	fN := float64(N)
	depth := c.Slices
	maxDepth := int(math.Ceil(fN / float64(c.SliceSize)))

	if maxDepth < depth {
		depth = maxDepth
	}
	depth--

	res := make([]byte, N+depth)
	off := 0
	for i := 0; i < N; i++ {
		if i != 0 && i%c.SliceSize == 0 && off < depth {
			res[i+off] = c.Sep
			off += 1
		}
		res[i+off] = text[i]
	}
	return string(res)
}
