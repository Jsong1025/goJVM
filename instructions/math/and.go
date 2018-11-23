package math

import (
	"goJVM/instructions/base"
	"goJVM/rtda"
)

type IAND struct {
	base.NoOperandsInstruction
}

type LAND struct {
	base.NoOperandsInstruction
}

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v2 & v1
	stack.PushInt(result)
}

func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v2 & v1
	stack.PushLong(result)
}