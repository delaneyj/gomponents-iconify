package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderMale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 5v1h3.3L13 10.3C12 9.5 10.8 9 9.5 9C6.5 9 4 11.5 4 14.5S6.5 20 9.5 20s5.5-2.5 5.5-5.5c0-1.3-.5-2.5-1.3-3.5L18 6.7V10h1V5h-5m-4.5 5c1 0 2 .4 2.8 1c.2.2.5.4.7.7c.6.8 1 1.8 1 2.8C14 17 12 19 9.5 19S5 17 5 14.5S7 10 9.5 10Z"/>`),
		g.Group(children),
	)
}