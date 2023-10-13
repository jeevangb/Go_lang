package temperature

func Temp(f float32) float32 {
	c := (f - 32) * 5 / 9
	return c
}
