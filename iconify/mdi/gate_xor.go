package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GateXor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4c3 6 3 10 0 16h2c3-6 3-10 .1-16H2m4 0c3 6 3 10 0 16h3c5 0 9-3 13-8c-4-5-8-8-13-8H6m3 2c3.8 0 7 2.1 10.3 6c-3.4 3.9-6.5 6-10.3 6c1.5-4 1.5-8 0-12Z"/>`),
		g.Group(children),
	)
}