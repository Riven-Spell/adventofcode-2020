package solutions

import (
	"sync"
	"sync/atomic"
)

type instName uint8

const (
	NOP instName = iota
	JMP
	ACC
)

var instructions = map[string]instName {
	"nop":NOP,
	"jmp":JMP,
	"acc":ACC,
}

type gcCpuInstruction struct {
	instName
	value int64
}

type gameConsoleCPU struct {
	Accumulator int64
	ExecutionHead int64
	Rom []gcCpuInstruction
}

func (cpu *gameConsoleCPU) Clone() gameConsoleCPU {
	gcCpu := gameConsoleCPU{
		Accumulator:   cpu.Accumulator,
		ExecutionHead: cpu.ExecutionHead,
		Rom:           make([]gcCpuInstruction, len(cpu.Rom)),
	}

	copy(gcCpu.Rom, cpu.Rom)

	return gcCpu
}

func (cpu *gameConsoleCPU) RunUntilTerminationOrForceHalt(haltBit *int64, wg *sync.WaitGroup) {
	defer wg.Done()

	ran := make(map[int64]bool)

	for int(cpu.ExecutionHead) < len(cpu.Rom) {
		ran[cpu.ExecutionHead] = true
		cpu.Step()

		// turns out Eric Wastl's only classification for the halting problem is that it doesn't hit the same code twice.
		// Fine by me.
		if _,ok := ran[cpu.ExecutionHead]; ok {
			return
		}

		if atomic.LoadInt64(haltBit) != 0 {
			return
		}
	}

	if int(cpu.ExecutionHead) != len(cpu.Rom) {
		return // invalid exit
	}

	// store the accumulator, intended for day 8 part 2
	atomic.StoreInt64(haltBit, cpu.Accumulator)
}

func (cpu *gameConsoleCPU) Step() {
	inst := cpu.Rom[cpu.ExecutionHead]

	switch inst.instName {
	case NOP:
		cpu.ExecutionHead++
		break
	case JMP:
		cpu.ExecutionHead += inst.value
		break
	case ACC:
		cpu.Accumulator += inst.value
		cpu.ExecutionHead++
		break
	}
}

func (cpu *gameConsoleCPU) Reset() {
	cpu.ExecutionHead = 0
	cpu.Accumulator = 0
}
