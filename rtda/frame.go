package rtda

import "goJVM/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPc       int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:		method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

//func NewFrame(maxLocals, maxStack uint) *Frame {
//	return &Frame{
//		localVars:    newLocalVars(maxLocals),
//		operandStack: newOperandStack(maxStack),
//	}
//}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) NextPc() int {
	return self.nextPc
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) SetNextPc(nextPC int) {
	self.nextPc = nextPC
}

func (self Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self Frame) OperandStack() *OperandStack {
	return self.operandStack
}
