package main

func BaseToBase(value string, base, newBase int) string {
	return DecToBase(BaseToDec(value, base), newBase)
}