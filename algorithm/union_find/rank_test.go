package union_find

import (
	"fmt"
	"testing"
)

func TestRankUnion(t *testing.T) {
	s := RankUnion[int]{}
	elements := []int{1, 2, 3}
	s.Init(elements)

	s.Merge(1, 2)
	s.Merge(1, 3)
	fmt.Println(s.parentsMap)
}
