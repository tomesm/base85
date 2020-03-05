package base85

import (
	"testing"
)

func TestEncode(t *testing.T) {

	base := base85.New()
	res1 := base.Encode("Donec nec euismod orci")

	if res1 != "6uQsS@j#Z#@j#?*Ble-0A0>f2@qb" {
		t.Errorf("want %q; got %q", "6uQsS@j#Z#@j#?*Ble-0A0>f2@qb", res1)
	}
}
