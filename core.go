// Copyright (c) 2015 Joshua Scoggins
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgement in the product documentation would be
//    appreciated but is not required.
// 2. Altered source versions must be plainly marked as such, and must not be
//    misrepresented as being the original software.
// 3. This notice may not be removed or altered from any source distribution.
//

// declaration of the iris2 core, a 64-bit variable length instruction vliw
// architecture with typed instructions
package iris2

type Core struct {
	//	Registers [RegisterCount]Register
	Memory []byte
}

func NewCore(memorySize uint) *Core {
	var c Core
	// initialize the registers
	/*
		for i := 0; i < RegisterCount; i++ {
			c.Registers[i] = 0
		}
	*/
	// construct the raw memory we are going to be using
	c.Memory = make([]byte, memorySize)
	return &c
}
