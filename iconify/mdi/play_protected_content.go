package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayProtectedContent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5v13h9v-2H4V7h13v4h2V5H2m7 4v5l3.5-2.5L9 9m12.04 2.67l-4.95 4.95l-2.13-2.12l-1.41 1.41l3.54 3.54l6.36-6.36l-1.41-1.42Z"/>`),
		g.Group(children),
	)
}