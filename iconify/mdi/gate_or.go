package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GateOr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4c3 6 3 10 0 16h6c5 0 11-4 14-8c-3-4-9-8-14-8H2m3 2h3c3.5 0 8.3 3 11.3 6c-3 3-7.8 6-11.3 6H5c1.4-4.1 1.4-7.9 0-12Z"/>`),
		g.Group(children),
	)
}