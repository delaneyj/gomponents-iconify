package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.41 8H17a1 1 0 0 0 0-2H7a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 6 7v10a1 1 0 0 0 2 0V9.41l8.29 8.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		g.Group(children),
	)
}