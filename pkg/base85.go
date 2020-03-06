package base85

type Base85 struct {
}

func New() *Base85 {
	return &Base85{}
}

func (b *Base85) Encode(text string) string {
	source := []rune(text)
	s := ""

	for len(source) > 0 {
		if len(source) < 4 {
			for i := len(source); i < 4; i++ {
				source = append(source, 0)
			}
		}
		s += b.encodeChunk(source[:4])
		source = source[4:]
	}
	return s
}

func (b *Base85) Decode(text string) string {
	source := []rune(text)
	s := ""

	for len(source) > 0 {
		s += b.decodeChunk(source[:5])
		source = source[5:]
	}
	//fmt.Printf("%s\n", s)
	return s
}

func (b *Base85) decodeChunk(chunk []rune) string {
	var decoded []rune
	var n uint64 = uint64(chunk[0] - 33)

	for i := 1; i < len(chunk); i++ {
		mod := uint64(chunk[i]) - 33
		n = (n * 85) + mod
	}
	for j := 0; j < 4; j++ {
		decoded = append(decoded, rune(n&0xFF))
		n = n >> 8
	}
	s := b.reverseRunes(string(decoded))
	return s
}

func (b *Base85) encodeChunk(chunk []rune) string {
	var encoded []rune
	var n uint64 = 0
	s := ""

	for _, v := range chunk {
		n = (n << 8) | uint64(v)
	}
	for i := 0; i < 5; i++ {
		mod := n % 85
		encoded = append(encoded, rune(mod+33))
		n = n / 85
	}
	for _, v := range encoded {
		s += string(v)
	}
	s = b.reverseRunes(s)
	return s
}

func (b *Base85) reverseRunes(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
