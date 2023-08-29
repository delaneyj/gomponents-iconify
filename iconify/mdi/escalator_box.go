package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3c-1.11 0-2 .89-2 2v14c0 1.11.89 2 2 2h14c1.11 0 2-.89 2-2V5c0-1.11-.89-2-2-2H5m9.34 3h2.71a2 2 0 0 1 2 2c0 1.11-.89 2-2 2H16l-6 8H7.05a2 2 0 1 1 0-4h1.29l6-8Z"/>`),
		g.Group(children),
	)
}