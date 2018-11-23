package refrences

import (
	"goJVM/instructions/base"
	"goJVM/rtda"
	"goJVM/rtda/heap"
)

type INVOKE_INTERFACE struct {
	index	uint
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader)  {
	self.index = uint(reader.ReadInt16())
	reader.ReadUint8()
	reader.ReadUint8()
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame)  {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolveMethod := methodRef.ResolvedInterfaceMethod()
	if resolveMethod.IsStatic() || resolveMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolveMethod.ArgSlotCount() -1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if ref.Class().IsImplements(methodRef.ResolveClass()) {
		panic("java.lang.IncompatibleChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}