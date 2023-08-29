package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinusSquareFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 13H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm-4 11H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z" opacity=".5"/>`),
		g.Group(children),
	)
}