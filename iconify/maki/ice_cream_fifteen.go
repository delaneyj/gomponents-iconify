package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCreamFifteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.44 8.17a3.48 3.48 0 0 0 2-.63c.464.323 1 .53 1.56.6h.44L8 13.7a.5.5 0 0 1-.92 0L5.44 8.17z" fill="currentColor"/><path d="M11.44 4.67a2 2 0 1 1-4 0a2 2 0 1 1-2-2h.12a2 2 0 1 1 3.75 0h.13a2 2 0 0 1 2 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}