; Fibonacci in Intel 8080 assembler.
; Results in b
start:
	xra	a	; zero out a
	mov	b, a	; b = a
	mov	c, a	; c = a
	adi	01h	; a = a + 1
	mov	c, a	; c = a
	xra	a	; zero out a
loop:	add	c	; a = a + c
	cmp	c
	jc	start	; jump if carry
	mov	b, a
	mov	a, c
	mov	c, b
	jmp	loop	; jump to loop