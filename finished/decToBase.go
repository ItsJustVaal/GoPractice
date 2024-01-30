package main

func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"
	var res string
	for dec > 0 {
		res += string(charset[dec%base])
		dec /= base
	}

	return Reverse(res)
}