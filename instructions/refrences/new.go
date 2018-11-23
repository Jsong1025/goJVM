package refrences

import (
	"goJVM/instructions/base"
	"goJVM/rtda"
	"goJVM/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classref := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classref.ResolveClass()

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
