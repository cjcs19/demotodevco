package conversor

import (
	"testing"
)

var TestCases map[int]string

func TestConvertRomanToNumeral(t *testing.T) {
	r := NewRoman()

	for key, value := range TestCases {
		n := r.ToNumber(value)
		if n != key {
			t.Errorf("I should return %v, but %v", key, n)
		}
		//fmt.Printf("Key %v retorno: %v value:%v\n", key, n, value)
	}

}

func TestConvertNumeralToRoman(t *testing.T) {
	r := NewRoman()

	for key, value := range TestCases {
		n := r.ToRoman(key)
		if n != value {
			t.Errorf("I should return %v, but %v", value, n)
		}
		//fmt.Printf("Key %v retorno: %v value:%v\n", key, n, value)
	}

}

func init() {

	TestCases = map[int]string{
		0:    "",
		1:    "I",
		2:    "II",
		4:    "IV",
		5:    "V",
		444:  "CDXLIV",
		555:  "DLV",
		666:  "DCLXVI",
		999:  "CMXCIX",
		1111: "MCXI",
		1993: "MCMXCIII",
		2018: "MMXVIII",
		2222: "MMCCXXII",
	}
}
