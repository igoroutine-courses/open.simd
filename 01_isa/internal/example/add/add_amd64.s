#include "textflag.h"

// func Add(a, b int64) int64
TEXT ·Add(SB), NOSPLIT, $0
    // Кладём первый аргумент функции в регистр AX
    MOVQ a+0(FP), AX

    // Кладём второй аргумент функции в регистр CX
    MOVQ b+8(FP), CX

    BYTE $0x48 // 01001000
    BYTE $0x01 // 00000001
    BYTE $0xc8 // 11001000

    // Кладём результат с AX на стек фрейм функции для возврата значения
    MOVQ AX, ret+16(FP)

    // Инструкция для return
    RET


// func Add(a, b int64) int64
// TEXT ·Add(SB), NOSPLIT, $0
    // BYTE $0x48
    // BYTE $0x8b
    // BYTE $0x44
    // BYTE $0x24
    // BYTE $0x08
    // BYTE $0x48
    // BYTE $0x8b
    // BYTE $0x4c
    // BYTE $0x24
    // BYTE $0x10
    // BYTE $0x48
    // BYTE $0x01
    // BYTE $0xc8
    // BYTE $0x48
    // BYTE $0x89
    // BYTE $0x44
    // BYTE $0x24
    // BYTE $0x18
    // BYTE $0xc3
