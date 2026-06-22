#include "textflag.h"

// func Sub(a, b int64) int64
TEXT ·Sub(SB), NOSPLIT, $0
    // Кладём первый аргумент функции в регистр R0
    MOVD a+0(FP), R0

    // Кладём второй аргумент функции в регистр R1
    MOVD b+8(FP), R1

    WORD $0xcb010001 // 11001011000000000000000000100001

    // Кладём результат с R1 на стек фрейм функции для возврата значения
    MOVD R1, ret+16(FP)

    // Инструкция для return
    RET

// func Sub(a, b int64) int64
// TEXT ·Sub(SB), NOSPLIT, $0
    // WORD $0xf94007e0
    // WORD $0xf9400be1
    // WORD $0xcb010001
    // WORD $0xf9000fe1
    // WORD $0xd65f03c0
