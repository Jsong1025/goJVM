package heap

import "goJVM/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func (self *Field) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Field) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Field) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Field) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Field) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Field) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Field) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}
func (self *Field) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	field := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		field[i] = &Field{}
		field[i].class = class
		field[i].copyMemberInfo(cfField)
		field[i].copyAttributes(cfField)
	}
	return field
}
