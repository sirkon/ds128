TEXT ·Mul64(SB), $0-40
    MOVD    lo+0(FP), R0
    MOVD    lo+16(FP), R1
	MUL     R0, R1, R2
	UMULH   R0, R1, R3
	MOVD    lo+8(FP), R1
	MUL     R0, R1, R4
	ADDS    R3, R4, R0
    MOVD.P  R2, ret+24(FP)
    MOVD.P  R0, ret+32(FP)
    RET
