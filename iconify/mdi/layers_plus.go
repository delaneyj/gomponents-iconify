package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 14h2v3h3v2h-3v3h-2v-3h-3v-2h3v-3m-6 2L2 9l9-7l9 7l-9 7m0 2.54l1-.79V18c0 .71.12 1.39.35 2L11 21.07l-9-7l1.62-1.26L11 18.54Z"/>`),
		g.Group(children),
	)
}