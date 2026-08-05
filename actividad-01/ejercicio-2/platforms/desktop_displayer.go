package platforms

import (
	"fmt"
)

type desktopDisplayer struct{}

func (dd desktopDisplayer) Display(kind string, message string) {
	message = fmt.Sprintf("[desktop] %s: %s", kind, message)
	fmt.Println(message)
}

func NewDesktopDisplayer() Displayer {
	return desktopDisplayer{}
}
