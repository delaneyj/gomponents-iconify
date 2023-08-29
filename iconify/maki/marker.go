package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 1C5.423 1 3 2.288 3 5.568C3 7.793 6.462 12.712 7.5 14c.923-1.288 4.5-6.09 4.5-8.432C12 2.288 9.577 1 7.5 1Z"/>`),
		g.Group(children),
	)
}