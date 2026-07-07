#include "textflag.h"

TEXT ·BroadcastInt32x4(SB), NOSPLIT, $0-20
    MOVW        value+0(FP), R0

    WORD        $0x4E040C00

    MOVD        $ret+4(FP), R1
    VST1        [V0.S4], (R1)
    RET


TEXT ·LoadInt32x4(SB), NOSPLIT, $0-24
    MOVD        ptr+0(FP), R0

    VLD1        (R0), [V0.S4]

    MOVD        $ret+8(FP), R1
    VST1        [V0.S4], (R1)
    RET


TEXT ·EqualInt32x4(SB), NOSPLIT, $0-48
    MOVD        $a+0(FP), R0
    VLD1        (R0), [V0.S4]

    MOVD        $b+16(FP), R1
    VLD1        (R1), [V1.S4]

    WORD        $0x6EA18C02

    MOVD        $ret+32(FP), R2
    VST1        [V2.S4], (R2)
    RET


TEXT ·MaskToLanes(SB), NOSPLIT, $0-24
    MOVD        $mask+0(FP), R0
    VLD1        (R0), [V0.S4]
    WORD        $0x6EB0A801           // UMAXV S1, V0.4S

    FMOVS       F1, R0

    MOVD        R0, ret+16(FP)
    RET


TEXT ·SliceContainsInt32(SB), NOSPLIT, $0-33
    MOVD        s_base+0(FP), R0
    MOVD        s_len+8(FP), R1
    MOVW        target+24(FP), R2
    MOVB        ZR, ret+32(FP)

    CBZ         R1, done

    WORD        $0x4E040C40           // DUP V0.4S, W2

    MOVD        R1, R3
    AND         $~3, R3, R3           // R3 = number of elements for SIMD loop
    MOVD        ZR, R4                // R4 = i = 0

simd_loop:
    CMP         R3, R4
    BGE         scalar_tail

    LSL         $2, R4, R5            // R5 = i * 4 (bytes offset)
    ADD         R0, R5, R5            // R5 = &s[i]
    VLD1        (R5), [V1.S4]

    WORD        $0x6EA08C22           // CMEQ V2.4S, V1.4S, V0.4S

    WORD        $0x6EB0A843           // UMAXV S3, V2.4S
    FMOVS       F3, R6
    CBNZ        R6, found

    ADD         $4, R4, R4            // i += 4
    B           simd_loop

scalar_tail:
    CMP         R1, R4
    BGE         done

    LSL         $2, R4, R5            // R5 = i * 4 (bytes offset)
    ADD         R0, R5, R5            // R5 = &s[i]
    MOVW        (R5), R6              // R6 = s[i]
    CMPW        R2, R6
    BEQ         found

    ADD         $1, R4, R4            // i++
    B           scalar_tail

found:
    MOVD        $1, R7
    MOVB        R7, ret+32(FP)

done:
    RET
