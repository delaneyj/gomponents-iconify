package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoeHeel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18h8.7l5.3-4h1v4h2v-4s1-2 1-4s-.5-4-.5-4h-2L18 7l-8 7H8l-5 2v2Z"/>`),
		g.Group(children),
	)
}