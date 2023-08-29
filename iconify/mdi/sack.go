package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 9c4 2 5 9 5 9s1 4-5 4H8c-6 0-5-4-5-4s1-7 5-9m6-5l-2-2l-2 2l-4-2l2 5h8l2-5l-4 2Z"/>`),
		g.Group(children),
	)
}