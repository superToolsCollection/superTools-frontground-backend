package util

import (
	"superTools-frontground-backend/pkg"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-24 10:04
* @Description:
**/

type Morse struct {
	str  string
	real string
	err  error
}

func TestGenerateMorse(t *testing.T) {
	var morses = []struct {
		str  string
		code string
		err  error
	}{
		{"aa11", ".-.-.----.----", nil},
		{"11aa", ".----.----.-.-", nil},
		{"", "", pkg.lengthError},
		{"111,as", "", pkg.unsupportedError},
		{"中文", "", pkg.unsupportedError},
		{"1a12 ", ".----.-.----..---", nil},
		{"   ", "", pkg.lengthError},
		{"asdj$%#, 441", "", pkg.unsupportedError},
		{"!@#$", "", pkg.unsupportedError},
	}

	for i, v := range morses {
		code, e := pkg.GenerateMorse(v.str)
		if code != v.code {
			t.Errorf("%d. %s morse code %s, wanted: %s, error= %v", i, v.str, code, v.code, e)
		} else if e != v.err {
			t.Errorf("%d. %s morse code %s, wanted: %s, error= %v", i, v.str, code, v.code, e)
		}
	}
}

func BenchmarkGenerateMorse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = pkg.GenerateMorse("asasd12454")
	}
}
