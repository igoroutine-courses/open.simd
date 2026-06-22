#include "textflag.h"

// func Min4(first, second, dst *[4]int32)
TEXT ·Min4(SB), NOSPLIT, $0
    // Кладём указатель на память first в регистр AX
    MOVQ first+0(FP), AX

    // Кладём указатель на память second в регистр CX
    MOVQ second+8(FP), CX

    // Кладём указатель на память dst в регистр DX
    MOVQ dst+16(FP), DX

    // Загружаем 128-бит в векторный регистр X1 (S - single precision, 4 элемента)

    // Загружаем 128-бит в векторный регистр X2 (S - single precision, 4 элемента)
    VMOVDQU (AX), X1
    VMOVDQU (CX), X2

    BYTE $0xc4 // 11000100
    BYTE $0xe2 // 11100010
    BYTE $0x71 // 01110001
    BYTE $0x39 // 00111001
    BYTE $0xda // 11011010

    // Записываем результат в память по указателю из регистра DX
    VMOVDQU X3, (DX)

    // Инструкция для return
    RET
