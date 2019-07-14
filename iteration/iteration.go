package iteration

func Repeat(seq string, numTimes int) (repeated string) {
	for i := 0; i < numTimes; i++ {
		repeated += seq
	}
	return
}
