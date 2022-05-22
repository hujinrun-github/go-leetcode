package union_find

import "fmt"

// CompressUnion -- the union find algorithm which has compressed the path
type CompressUnion[T int | ~string | float32] struct {
	parentsMap map[T]T
}

// Init -- init simple union find
func (s *CompressUnion[T]) Init(initList []T) error {
	if s.parentsMap == nil {
		s.parentsMap = map[T]T{}
	}
	if len(initList) != 0 {
		for _, item := range initList {
			s.parentsMap[item] = item
		}
	} else {
		return fmt.Errorf("init union_find with nil")
	}
	return nil
}

// Find -- find the parent of 'value'
func (s *CompressUnion[T]) Find(value T) T {
	if s.parentsMap[value] == value {
		return value
	} else {
		// compress the path
		s.parentsMap[value] = s.Find(s.parentsMap[value])
		return s.Find(s.parentsMap[value])
	}
}

// Merge -- merge union 'i' and 'j', change the parent of 'i' to the parent of 'j'
func (s *CompressUnion[T]) Merge(i T, j T) {
	s.parentsMap[s.Find(i)] = s.Find(j)
}
