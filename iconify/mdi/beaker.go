package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3v2a2 2 0 0 1 2 2v12a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V7a2 2 0 0 1 2-2V3H3m4 6h3v1H7V9m0 2h3v1H7v-1m3 5H7v-1h3v1m2-2H7v-1h5v1m0-6H7V7h5v1Z"/>`),
		g.Group(children),
	)
}