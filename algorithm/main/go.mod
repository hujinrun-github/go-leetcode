module main

go 1.18

replace union_find => ../union_find

replace graph/base => ../graph/base

require (
	graph/base v0.0.0-00010101000000-000000000000 // indirect
	union_find v0.0.0-00010101000000-000000000000 // indirect
)
