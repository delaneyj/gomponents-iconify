package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 10H4V6h11a1 1 0 0 1 1 1v3z" opacity=".5"/><path fill="currentColor" d="M21 18H4v-8h17a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1z"/><path fill="currentColor" d="M3 22a1 1 0 0 1-1-.999V3a1 1 0 0 1 2 0v18a1 1 0 0 1-.999 1H3z" opacity=".25"/>`),
		g.Group(children),
	)
}