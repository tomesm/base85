package base85

type Base85 struct {
}

func New() *Base85 {
	return &Base85{}
}

func (b *Base85) Encode(text string) string {
	source := []rune(text)
	str := ""
	cut := 0

	for len(source) > 0 {
		if len(source) < 4 {
			cut = len(source) + 1

			for i := len(source); i < 4; i++ {
				source = append(source, 0)
			}
		}
		str += b.encodeChunk(source[:4], cut)
		source = source[4:]
	}
	return str
}

func (b *Base85) encodeChunk(chunk []rune, cut int) string {
	var encoded []rune
	var n uint64 = 0
	str := ""

	for _, v := range chunk {
		n = (n << 8) | uint64(v)
	}
	for i := 0; i < 5; i++ {
		mod := n % 85
		encoded = append(encoded, rune(mod+33))
		n = n / 85
	}
	for _, v := range encoded {
		str += string(v)
	}
	str = b.reverseString(str)

	if cut != 0 && cut < 4 {
		str = str[:cut]
	}
	return str
}

func (b *Base85) reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
