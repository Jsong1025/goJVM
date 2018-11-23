package rtda

import "goJVM/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
