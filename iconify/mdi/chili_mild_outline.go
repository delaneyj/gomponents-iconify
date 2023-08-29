package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChiliMildOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.25 7.5l-1.52-.87c.53-.93 1.3-1.63 2.21-1.94c-.14-.4-.5-.69-.94-.69V2c1.54 0 2.79 1.16 2.97 2.65c.94.29 1.75 1.01 2.3 1.98l-1.52.87l-1.75-1l-1.75 1M16 10v12s-8-2-8-11v-1c0-.73.4-1.37 1-1.72l1.25.72L12 8l1.75 1L15 8.28c.6.35 1 .99 1 1.72m-2 1.45l-2-1.15l-2 1.13c.17 4.01 2.23 6.26 4 7.44v-7.42Z"/>`),
		g.Group(children),
	)
}