package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAmountDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.29 14.29l-.29.3V7a1 1 0 0 0-2 0v7.59l-.29-.3a1 1 0 0 0-1.42 1.42l2 2a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2-2a1 1 0 0 0-1.42-1.42ZM11 8h10a1 1 0 0 0 0-2H11a1 1 0 0 0 0 2Zm10 3H11a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2Zm0 5H11a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}