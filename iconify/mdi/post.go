package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Post(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3v18h18V3H3m15 15H6v-1h12v1m0-2H6v-1h12v1m0-4H6V6h12v6Z"/>`),
		g.Group(children),
	)
}