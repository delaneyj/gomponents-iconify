package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListChecks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5h6v6H3V5m2 2v2h2V7H5m6 0h10v2H11V7m0 8h10v2H11v-2m-6 5l-3.5-3.5l1.41-1.41L5 17.17l4.59-4.58L11 14l-6 6Z"/>`),
		g.Group(children),
	)
}