package conversions

import (
	"goJVM/instructions/base"
	"goJVM/rtda"
)

type F2F struct {
	base.NoOperandsInstruction
}

type F2I struct {
	base.NoOperandsInstruction
}

type F2L struct {
	base.NoOperandsInstruction
}

type F2D struct {
	base.NoOperandsInstruction
}

func (self *F2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := float32(d)
	stack.PushFloat(i)
}

func (self *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := int32(d)
	stack.PushInt(i)
}

func (self *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := int64(d)
	stack.PushLong(i)
}

func (self *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := float64(d)
	stack.PushDouble(i)
}