package stringutil_test

import (
	"reflect"
	"testing"

	"github.com/naoina/go-stringutil"
)

func TestToSnakeCase(t *testing.T) {
	for _, v := range []struct {
		input, expect string
	}{
		{"", ""},
		{"thequickbrownfoxjumpsoverthelazydog", "thequickbrownfoxjumpsoverthelazydog"},
		{"Thequickbrownfoxjumpsoverthelazydog", "thequickbrownfoxjumpsoverthelazydog"},
		{"ThequickbrownfoxjumpsoverthelazydoG", "thequickbrownfoxjumpsoverthelazydo_g"},
		{"TheQuickBrownFoxJumpsOverTheLazyDog", "the_quick_brown_fox_jumps_over_the_lazy_dog"},
		{"the_quick_brown_fox_jumps_over_the_lazy_dog", "the_quick_brown_fox_jumps_over_the_lazy_dog"},
		{"ＴｈｅＱｕｉｃｋＢｒｏｗｎＦｏｘＯｖｅｒＴｈｅＬａｚｙＤｏｇ", "ｔｈｅ_ｑｕｉｃｋ_ｂｒｏｗｎ_ｆｏｘ_ｏｖｅｒ_ｔｈｅ_ｌａｚｙ_ｄｏｇ"},
	} {
		actual := stringutil.ToSnakeCase(v.input)
		expect := v.expect
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf(`stringutil.ToSnakeCase(%#v) => %#v; want %#v`, v.input, actual, expect)
		}
	}
}
