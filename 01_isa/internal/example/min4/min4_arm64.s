#include "textflag.h"

// func Min4(first, second, dst *[4]int32)
TEXT ·Min4(SB), NOSPLIT, $0
    // Кладём указатель на память first в регистр R0
    MOVD first+0(FP), R0

    // Кладём указатель на память second в регистр R1
    MOVD second+8(FP), R1

    // Кладём указатель на память dst в регистр R2
    MOVD dst+16(FP), R2

    // Загружаем 128-бит в векторный регистр V1 (S - single precision, 4 элемента)
    VLD1 (R0), [V1.S4]

    // Загружаем 128-бит в векторный регистр V2 (S - single precision, 4 элемента)
    VLD1 (R1), [V2.S4]

    WORD $0x4ea26c23 // 01001110101000100110110000100011

    // Записываем результат в память по указателю из регистра R2
    VST1 [V3.S4], (R2)

    // Инструкция для return
    RET
