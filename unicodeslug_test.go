package unicodeslug

import (
	"regexp"
	"testing"
)

var tests []string = []string{
	" 파이썬-기초강좌-1",
	"파이썬-기초강좌-1 ",
	"파이썬 기초강좌-1",
	"파이썬-기초강좌-!1",
	"파이썬-기초강좌-@1",
	"파이썬-기초강좌-#1",
	"파이썬-기초강좌-$1",
	"파이썬-기초강좌-%1",
	"파이썬-기초강좌-^1",
	"파이썬-기초강좌-&1",
	"파이썬-기초강좌-*1",
	"파이썬-기초강좌-(1",
	"파이썬-기초강좌-)1",
	"파이썬-기초강좌-/1",
	"파이썬-기초강좌-\\1",
	"파이썬-기초강좌-1!",
	"파이썬-기초강좌-1@",
	"파이썬-기초강좌-1#",
	"파이썬-기초강좌-1$",
	"파이썬-기초강좌-1%",
	"파이썬-기초강좌-1^",
	"파이썬-기초강좌-1&",
	"파이썬-기초강좌-1*",
	"파이썬-기초강좌-1(",
	"파이썬-기초강좌-1)",
	"파이썬-기초강좌-1/",
	"파이썬-기초강좌-1\\",
	"-파이썬-기초강좌-1",
	"파이썬-기초강좌-1-",
}

func TestSlugify(t *testing.T) {
	for _, test := range tests {
		slug := Slugify(test)
		hasSpecialCharacter(t, slug)
	}
}

func TestValidateSlug(t *testing.T) {
	for _, test := range tests {
		if ValidateSlug(test) {
			t.Errorf("Test ValidateSlug failed: %s", test)
		}
	}
}

func hasSpecialCharacter(t *testing.T, str string) {
	re := regexp.MustCompile(`[!@#$%^*/\s]|^-|-$`)
	match := re.FindString(str)
	if len(match) > 0 {
		t.Errorf("\"%s\" has a special character \"%s\"", str, match)
	}
}
