package i18n

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	Location = "locale/"
	Language = "ru"

	r, err := T("error.zero.none")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(r)
}
