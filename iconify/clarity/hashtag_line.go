package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashtagLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path id="clarityHashtagLine0" fill="currentColor" d="M32 12h-6.66l1.55-7.74a1 1 0 0 0-2-.39L23.3 12h-8.19l1.55-7.74a1 1 0 0 0-2-.39L13.07 12H6a1 1 0 0 0 0 2h6.67l-1.6 8H4a1 1 0 0 0 0 2h6.66l-1.55 7.74a1 1 0 0 0 .79 1.17a.68.68 0 0 0 .2 0a1 1 0 0 0 1-.8L12.7 24h8.19l-1.55 7.74a1 1 0 0 0 .79 1.17a.62.62 0 0 0 .19 0a1 1 0 0 0 1-.8L22.93 24H30a1 1 0 0 0 0-2h-6.67l1.61-8H32a1 1 0 0 0 0-2ZM21.29 22H13.1l1.61-8h8.19Z"/>`),
		g.Group(children),
	)
}