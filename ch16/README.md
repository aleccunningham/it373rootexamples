# Segmentation

- Instead of having just one base/bounds pair in the MMU, create a base and bounds pair per logical segments of the address space
- A segment is just a contiguous portion of the address space of a particular length
- Segmentation Fault - violation that arises from a memory access on a segmented machine to an illegal address

**Segment Registers:**

Segment|Base|Size|Grows Positive|Protection
-------|--- |----|--------------|---------
Code   |32k |2K  |1             |Read-Excecute
Heap   |34k |2K  |1             |Read-Write
Stack  |28k |2K  |0             |Read-Write 
