package platforms

type Displayer interface {
	Display(kind string, message string)
}
