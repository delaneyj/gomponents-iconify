package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTableBoxMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5v16h16v2H3c-1.1 0-2-.9-2-2V5h2m18-4H7c-1.11 0-2 .89-2 2v14a2 2 0 0 0 2 2h14c1.11 0 2-.89 2-2V3a2 2 0 0 0-2-2M11 16H8v-2h3v2m0-3H8v-2h3v2m0-3H8V8h3v2m4 6h-3v-2h3v2m0-3h-3v-2h3v2m0-3h-3V8h3v2Z"/>`),
		g.Group(children),
	)
}