package heap

import (
	"fmt"
	"goJVM/classfile"
)

type Constant interface {
}

type ConstantPool struct {
	class	*Class
	consts	[]Constant
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()

		case *classfile.ConstantFloatInfo:
			intInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = intInfo.Value()

		case *classfile.ConstantLongInfo:
			intInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = intInfo.Value()
			i++

		case *classfile.ConstantDoubleInfo:
			intInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = intInfo.Value()
			i++

		case *classfile.ConstantStringInfo:
			intInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = intInfo.String()
			i++

		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)

		case *classfile.ConstantFieldInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)

		case *classfile.ConstantMethodRefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)

		case *classfile.ConstantInterfaceMethodRefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodRefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)

		}

	}

	return rtCp
}