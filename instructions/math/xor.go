package math

import (
	"goJVM/instructions/base"
	"goJVM/rtda"
)

type IXOR struct {
	base.NoOperandsInstruction
}

type LXOR struct {
	base.NoOperandsInstruction
}

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v2 ^ v1
	stack.PushInt(result)
}

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v2 ^ v1
	stack.PushLong(result)
}