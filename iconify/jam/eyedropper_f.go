package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyedropperF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.473 11.693l-1.36 1.36l-8.417-.068L9.23 7.451a1 1 0 1 1 1.415-1.414l4.242-4.243a3 3 0 1 1 4.243 4.243l-4.243 4.242a1 1 0 0 1-1.414 1.414zm-3.191 3.192L7.57 17.596a7 7 0 0 1-2.736 1.691l-2.899.966A1 1 0 0 1 .67 18.988l.967-2.899a7 7 0 0 1 .522-1.186l8.123-.018z"/>`),
		g.Group(children),
	)
}