package base85

import (
	"testing"
)

func TestEncode(t *testing.T) {
	base := NewCoder()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "With short input",
			input: "Man",
			want:  "9jqo>",
		},
		{
			name:  "With long input",
			input: "Donec nec euismod orci",
			want:  "6uQsS@j#Z#@j#?*Ble-0A0>f2@qboC",
		},
		{
			name:  "With complex input",
			input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
			want:  `9Q+r_D'3P3F*2=BA8c:&EZfF;F<G"/ATTIG@rH7+ARfgnFEMUH@:X(kBldcuDJ()'Ch[tk`,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			res := base.Encode([]rune(testCase.input))

			if res != testCase.want {
				t.Errorf("want %q; got %q", testCase.want, res)
			}
		})
	}
}
