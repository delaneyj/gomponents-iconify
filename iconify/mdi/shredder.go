package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shredder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3v4h2V5h8v2h2V3H6M5 8a3 3 0 0 0-3 3v6h3v-3h14v3h3v-6a3 3 0 0 0-3-3H5m13 2a1 1 0 0 1 1 1a1 1 0 0 1-1 1a1 1 0 0 1-1-1a1 1 0 0 1 1-1M7 16v5h2v-5H7m4 0v4h2v-4h-2m4 0v5h2v-5h-2Z"/>`),
		g.Group(children),
	)
}