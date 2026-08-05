package platforms

import (
	"fmt"
)

type webDisplayer struct{}

func (wd webDisplayer) Display(kind string, message string) {
	message = fmt.Sprintf("[web] %s: %s", kind, message)
	fmt.Println(message)
}

func NewWebDisplayer() Displayer {
	return webDisplayer{}
}
