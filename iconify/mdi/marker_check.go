package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkerCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 16l-5-5l1.41-1.42L10 13.17l7.59-7.59L19 7m0-6H5c-1.11 0-2 .89-2 2v12.93c0 .69.35 1.3.88 1.66L12 23l8.11-5.41c.53-.36.89-.97.89-1.66V3a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}