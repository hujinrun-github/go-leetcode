package union_find

import "fmt"

type unionInfo[T int | ~string | float32] struct {
	parent T
	rank   int
}

// RankUnion -- the union find algorithm which has compressed the path use rank less principle
type RankUnion[T int | ~string | float32] struct {
	parentsMap map[T]unionInfo[T]
}

// Init -- init simple union find
func (r *RankUnion[T]) Init(initList []T) error {
	if r.parentsMap == nil {
		r.parentsMap = map[T]unionInfo[T]{}
	}
	if len(initList) != 0 {
		for _, item := range initList {
			r.parentsMap[item] = unionInfo[T]{
				parent: item,
				rank:   1,
			}
		}
	} else {
		return fmt.Errorf("init union_find with nil")
	}
	return nil
}

// Find -- find the parent of 'value'
func (r *RankUnion[T]) Find(value T) T {
	if r.parentsMap[value].parent == value {
		return value
	} else {
		// compress the path
		//r.parentsMap[value].parent = r.Find(r.parentsMap[value].parent)
		r.parentsMap[value] = unionInfo[T]{
			parent: r.Find(r.parentsMap[value].parent),
			rank:   r.parentsMap[value].rank,
		}
		return r.Find(r.parentsMap[value].parent)
	}
}

// Merge -- merge union 'i' and 'j', change the parent of 'i' to the parent of 'j'
func (r *RankUnion[T]) Merge(i T, j T) {
	iparent := r.parentsMap[i].parent
	jparent := r.parentsMap[j].parent
	if r.parentsMap[iparent].rank <= r.parentsMap[jparent].rank {
		r.parentsMap[i] = unionInfo[T]{
			parent: jparent,
			rank:   r.parentsMap[i].rank,
		}
	} else {
		r.parentsMap[j] = unionInfo[T]{
			parent: iparent,
			rank:   r.parentsMap[j].rank,
		}
	}

	if r.parentsMap[iparent].rank == r.parentsMap[jparent].rank && iparent != jparent {
		r.parentsMap[j] = unionInfo[T]{
			parent: r.parentsMap[j].parent,
			rank:   r.parentsMap[j].rank + 1,
		}
	}
}
