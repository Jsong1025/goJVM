package constants

import (
	"goJVM/instructions/base"
	"goJVM/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么也不做
}
