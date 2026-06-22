#include "textflag.h"

// func Add(a, b int64) int64
TEXT ·Add(SB), NOSPLIT, $0
    // Кладём первый аргумент функции в регистр R0
    MOVD a+0(FP), R0

    // Кладём второй аргумент функции в регистр R1
    MOVD b+8(FP), R1

    WORD $0x8b000021 // 10001011000000000000000000100001

    // Кладём результат с R1 на стек фрейм функции для возврата значения
    MOVD R1, ret+16(FP)

    // Инструкция для return
    RET

// func Add(a, b int64) int64
// TEXT ·Add(SB), NOSPLIT, $0
    // WORD $0xf94007e0 // MOVD 8(RSP), R0 | MOVD first+0(FP), R0
    // WORD $0xf9400be1 // MOVD 16(RSP), R1 | MOVD second+8(FP), R1
    // WORD $0x8b000021 // ADD R0, R1, R1
    // WORD $0xf9000fe1 // MOVD R1, 24(RSP) | MOVD R1, ret+16(FP)
    // WORD $0xd65f03c0 // RET
