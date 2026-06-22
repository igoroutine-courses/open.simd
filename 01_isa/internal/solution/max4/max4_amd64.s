#include "textflag.h"

// func Max4(first, second, dst *[4]int32)
TEXT ·Max4(SB), NOSPLIT, $0

    MOVQ first+0(FP), AX
    MOVQ second+8(FP), CX
    MOVQ dst+16(FP), DX

    VMOVDQU (AX), X1
    VMOVDQU (CX), X2

    // X3[i] = MAX(X1[i], X2[i]) (VPMAXSD)
    BYTE $0xc4
    BYTE $0xe2
    BYTE $0x71
    BYTE $0x3d
    BYTE $0xda

    VMOVDQU X3, (DX)
    RET
