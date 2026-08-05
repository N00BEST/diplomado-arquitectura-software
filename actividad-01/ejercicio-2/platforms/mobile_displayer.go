package platforms

import (
	"fmt"
)

type mobileDisplayer struct{}

func (md mobileDisplayer) Display(kind string, message string) {
	message = fmt.Sprintf("[mobile] %s: %s", kind, message)
	fmt.Println(message)
}

func NewMobileDisplayer() Displayer {
	return mobileDisplayer{}
}
