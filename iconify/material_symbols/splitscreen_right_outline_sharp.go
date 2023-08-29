package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitscreenRightOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21V3h8v18h-8ZM3 21V3h8v18H3ZM5 5v14h4V5H5Zm4 14H5h4Z"/>`),
		g.Group(children),
	)
}