package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeBlockBraces(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3c-1.1 0-2 .9-2 2s-.9 2-2 2v2c1.1 0 2 .9 2 2s.9 2 2 2h2v-2H5v-1c0-1.1-.9-2-2-2c1.1 0 2-.9 2-2V5h2V3m4 0c1.1 0 2 .9 2 2s.9 2 2 2v2c-1.1 0-2 .9-2 2s-.9 2-2 2H9v-2h2v-1c0-1.1.9-2 2-2c-1.1 0-2-.9-2-2V5H9V3h2m11 3v12c0 1.11-.89 2-2 2H4a2 2 0 0 1-2-2v-3h2v3h16V6h-2.97V4H20c1.11 0 2 .89 2 2Z"/>`),
		g.Group(children),
	)
}