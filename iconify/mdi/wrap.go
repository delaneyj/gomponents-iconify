package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 5H3v2h18V5M3 19h7v-2H3v2m0-6h15c1 0 2 .43 2 2s-1 2-2 2h-2v-2l-4 3l4 3v-2h2c2.95 0 4-1.27 4-4c0-2.72-1-4-4-4H3v2Z"/>`),
		g.Group(children),
	)
}