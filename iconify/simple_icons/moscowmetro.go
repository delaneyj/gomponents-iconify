package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moscowmetro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.603 11.85l-2.481-6.26l-2.092 3.66l-2.092-3.66l-2.481 6.262H6.74v.941h3.736v-.941h-.553l.538-1.555l1.569 2.57l1.569-2.57l.538 1.555h-.553v.941h3.751v-.941zm5.335-1.912A9.933 9.933 0 0 0 12 0C6.516 0 2.062 4.453 2.062 9.938c0 2.75 1.121 5.23 2.914 7.023a.804.804 0 0 0 1.375-.568a.825.825 0 0 0-.239-.582a8.303 8.303 0 0 1-2.42-5.873c0-4.588 3.72-8.324 8.308-8.324c4.588 0 8.324 3.736 8.324 8.324a8.289 8.289 0 0 1-2.436 5.888l-7.024 7.023L12 24l7.039-7.039a9.891 9.891 0 0 0 2.899-7.023Z"/>`),
		g.Group(children),
	)
}