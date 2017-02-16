package gotools

func byte_inverse(nDauVao byte) byte {
	var x int = 1
	var y int = 0
	var u int = 0
	var v int = 1
	var q int
	var r int
	var s int
	var t int
	var a int = 6
	var b int = int(nDauVao)
	for b != 0 {
		q = a / b
		r = a - b * q
		s = x - u * q
		t = y - v * q
		a = b
		x = u
		y = v
		b = r
		u = s
		v = t
	}
	for y < 0 {
		y += 256
	}
	for y > 256 {
		y -= 256
	}
	return byte(y)
}