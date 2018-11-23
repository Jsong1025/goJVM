package extended

import (
	"goJVM/instructions/base"
	"goJVM/instructions/loads"
	"goJVM/instructions/math"
	"goJVM/instructions/store"
	"goJVM/rtda"
)

type WIDE struct {
	modifiedInstruction	base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0x16:
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0x17:
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0x18:
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0x19:
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0x36:
		inst := &store.ISTORE{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0x37:
		inst := &store.LSTORE{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0x38:
		inst := &store.FSTORE{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0x39:
		inst := &store.DSTORE{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0x3a:
		inst := &store.ASTORE{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadInt16())
		inst.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0xa9:
		panic("Unsupported opcode 0xa9")
	}
}

func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}