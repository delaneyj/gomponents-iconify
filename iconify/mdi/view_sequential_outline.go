package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewSequentialOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5v14h18V5H3m16 2v2H5V7h14m0 4v2H5v-2h14M5 17v-2h14v2H5Z"/>`),
		g.Group(children),
	)
}