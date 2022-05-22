package union_find

import (
	"fmt"
	"testing"
)

func TestCompressUnion(t *testing.T) {
	c := CompressUnion[int]{}
	elements := []int{1, 2, 3, 4, 5}
	c.Init(elements)

	c.Merge(1, 2)
	c.Merge(1, 3)
	c.Merge(1, 4)
	c.Merge(1, 5)

	fmt.Println(c.parentsMap)
}
