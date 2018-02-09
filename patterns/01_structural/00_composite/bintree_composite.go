// Binary Tree Compositions
//
// Is a approach to the composite pattern by storing instances of itself in a
// field. Working as a recursive compositing, and, because of the nature of
// recursivity, we must use pointers so that the compiler knows how much memory
// it must reserve for this struct. Our Tree struct stored a LeafValue object
// for each instance and a new Tree in its Right and Left fields.
//
package main

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

func main() {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
			Left:      nil,
		},
		Left: &Tree{4, nil, nil},
	}

	println(root.Right.Right.LeafValue)
}
