package nishan

import (
	"testing"
)

func Test_isCnicValid(t *testing.T) {
	valid := []string{"1513231551513", "3425321567541"}
	invalid := []string{"151323155151a", "342532156754", "34253215675412", "3425321567541a"}

	check(t, isCnicValid, valid, invalid)
}

func Test_isImeiValid(t *testing.T) {
	valid := []string{"151322132551513", "342532156754153"}
	invalid := []string{"151323155151a12", "34253215675423", "3425321567541213", "3425321567541a12"}

	check(t, isImeiValid, valid, invalid)
}

func Test_isMobileNumberValid(t *testing.T) {
	valid := []string{
		"+923001234567", "+92 3001234567", "+92 300 1234567",
		"923001234567", "92 3001234567", "92 300 1234567",
		"03001234567", "0300 1234567",
		"3001234567", "300 1234567",
	}
	invalid := []string{
		"+13001234567", "+92 400 1234567",
		"0370 1234567", "0 300 1234567",
		"1300 1234567", "0400 1234567",
		"4001234567", "400 1234567",
		"+92 300 12345678", "9230012345678",
		"0300 12345678", "030012345678", "30012345678",
		"O3OO 1234567",
	}

	check(t, isMobileNumberValid, valid, invalid)
}

func Test_isNumeric(t *testing.T) {
	valid := []string{"1234", "0", ""}
	invalid := []string{"123a", "O"}

	check(t, isNumeric, valid, invalid)
}

func Test_isBase64Encoded(t *testing.T) {
	valid := []string{"aGk=", "dGVzdA==", "YW4gM3h0ZXJNZWx5IGxvbmcgc3RyaW5nICQlIyUoXiMhKA=="}
	invalid := []string{"aG k=", "aG%k="}

	check(t, isBase64Encoded, valid, invalid)
}

func check(t *testing.T, f func(string) bool, valid, invalid []string) {
	for _, val := range valid {
		if !f(val) {
			t.Errorf("`%s` should be valid", val)
		}
	}

	for _, val := range invalid {
		if f(val) {
			t.Errorf("`%s` should be invalid", val)
		}
	}
}

func Test_isLatitudeValid(t *testing.T) {
	valid := []float64{-90, -45, 0, 45, 90}
	invalid := []float64{-180, -91, -90.001, 90.001, 91, 180}

	checkFloat(t, isLatitudeValid, valid, invalid)
}

func Test_isLongitudeValid(t *testing.T) {
	valid := []float64{-180, -90, 0, 45, 90, 180}
	invalid := []float64{-360, -181, -180.001, 180.001, 181, 360}

	checkFloat(t, isLongitudeValid, valid, invalid)
}

func checkFloat(t *testing.T, f func(float64) bool, valid, invalid []float64) {
	for _, val := range valid {
		if !f(val) {
			t.Errorf("`%f` should be valid", val)
		}
	}

	for _, val := range invalid {
		if f(val) {
			t.Errorf("`%f` should be invalid", val)
		}
	}
}
