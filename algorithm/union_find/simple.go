package union_find

type SimpleUnion[T int | ~string] struct {
	ParentsMap map[T]T
}
