package heap

import "goJVM/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func (self *Method) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Method) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Method) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Method) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Method) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Method) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Method) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}
func (self *Method) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}

func (self *Method) calcArgSlotCount() {
	parseDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parseDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}
}