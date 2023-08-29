package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 6v6H7.83l2.58-2.59L9 8l-5 5l5 5l1.41-1.41L7.83 14H20V6h-2Z"/>`),
		g.Group(children),
	)
}