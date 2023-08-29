package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionsBookmarkSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18V2h16v16H6Zm-4 4V6h2v14h14v2H2ZM13 4v7l2.5-1.5L18 11V4h-5Z"/>`),
		g.Group(children),
	)
}