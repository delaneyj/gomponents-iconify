package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 5.75L10.25 10H7v6h6.5l2 2H7a2 2 0 0 1-2-2v-6H1.75L6 5.75m12 12.5L13.75 14H17V8h-6.5l-2-2H17a2 2 0 0 1 2 2v6h3.25L18 18.25Z"/>`),
		g.Group(children),
	)
}