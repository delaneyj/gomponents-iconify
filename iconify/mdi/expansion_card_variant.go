package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpansionCardVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7h2.5v10H3V8.5H2M22 7v9h-8v1H7v-1H6V7m4 2H8v3h2m3-3h-2v3h2m7-3h-5v5h5V9Z"/>`),
		g.Group(children),
	)
}