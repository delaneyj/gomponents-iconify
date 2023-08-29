package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 7H7v4h4V7Zm0 6H7v4h4v-4Zm2 0h4v4h-4v-4Zm4-6h-4v4h4V7Z"/>`),
		g.Group(children),
	)
}