package randstr

var (
	stringStream = NewStream("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numberStream = NewStream("012345678901234567890123456789") // do not shorten this
)

func String(n uint) string {
	return stringStream.Next(n)
}

func Number(n uint) string {
	return numberStream.Next(n)
}
