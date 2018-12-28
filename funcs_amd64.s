TEXT ·Mul64(SB), $0-40
    MOVQ    lo+0(FP), AX
    MOVQ    lo+16(FP), BX
    MULQ    BX
    MOVQ    AX, ret+24(FP)
    MOVQ    lo+8(FP), AX
    MOVQ    DX, CX
    MULQ    BX
    ADDQ    CX, AX
    MOVQ    AX, ret+32(FP)
    RET


