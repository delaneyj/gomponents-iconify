package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5.44 8.17a3.473 3.473 0 0 0 2-.63a3.5 3.5 0 0 0 1.56.6h.44L8 13.7a.5.5 0 0 1-.92 0Zm6-3.473a2 2 0 0 1-4 0a2 2 0 1 1-2-2h.12a2 2 0 1 1 3.75 0h.13A2 2 0 0 1 11.439 4.7Z"/>`),
		g.Group(children),
	)
}