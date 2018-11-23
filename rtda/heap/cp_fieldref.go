package heap

import "goJVM/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (self *FieldRef) ResolveField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolveClass()
	field := lookupFiled(c, self.name, self.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

func lookupFiled(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := lookupFiled(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return lookupFiled(c.superClass, name, descriptor)
	}
	return nil
}
