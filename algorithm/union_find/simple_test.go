package union_find

import (
	"fmt"
	"testing"
)

func TestSimpleUnion(t *testing.T) {
	s := SimpleUnion[int]{}
	elements := []int{1, 2, 3, 4, 5, 6}
	s.Init(elements)

	s.Merge(2, 1)
	s.Merge(3, 2)
	s.Merge(5, 4)
	s.Merge(6, 5)

	if s.Find(2) != 1 {
		t.Errorf("the real parents of 2 is :%v", s.Find(2))
	}

	if s.Find(3) != 1 {
		t.Errorf("the real parantes of 3 is :%v", s.Find(3))
	}

	if s.Find(5) != 4 {
		t.Errorf("the real parantes of 5 is :%v", s.Find(3))
	}

	if s.Find(6) != 4 {
		t.Errorf("the real parantes of 6 is :%v", s.Find(3))
	}

	fmt.Println(s.parentsMap)
}
