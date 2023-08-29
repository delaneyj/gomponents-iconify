package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 2.5l8.84 3.15l-1.34 11.7L12 21.5l-7.5-4.15l-1.34-11.7L12 2.5m0 2.1L6.47 17h2.06l1.11-2.78h4.7L15.45 17h2.05L12 4.6m1.62 7.9h-3.23L12 8.63l1.62 3.87Z"/>`),
		g.Group(children),
	)
}