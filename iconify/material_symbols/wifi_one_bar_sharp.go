package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOneBarSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 21l-3.525-3.55q.7-.7 1.613-1.075T12 16q1 0 1.913.375t1.612 1.075L12 21Z"/>`),
		g.Group(children),
	)
}