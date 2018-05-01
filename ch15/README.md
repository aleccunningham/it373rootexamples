# Mechanism: Address Translation

- Limited direct execution
  - Let the program run directly on the hardware; however at certain key points in time, arrange so that the OS gets involved and makes sure the "right" thing happens.
- Hardware-based address translation (Address translation)
  - Hardware transforms each memory access request, changing the virtual address to the physical address where the desired information is already located.
- Memory Management Unit (MMU) is the part of the processor that helps translate virtual to physical addresses
- Dynamic Relocation
  - "base and bounds" (registers that are part of the hardware)
  - Each program is written and compiled as if it is loaded at address zero
  - The OS will load the process somewhere in the physical memory and returns the address
  - `physical addres = virtual address + base`
  - Bounds register is used to ensure safety; the processor will first check that a memory reference is *within bounds* to make sure it's legal
  - Free Lists
    - Primitive data structure for storing a list of currently free blocks of memory
  - When the OS decides to stop running a process, it must save the values of the base and bounds registers to memory, in some per-process structure such as the process structure or process control block (PCB) 
