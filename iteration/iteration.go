package iteration

func Repeat(toRepeat string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += toRepeat
	}
	return
}
