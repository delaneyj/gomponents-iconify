package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionsBookmarkOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16h12V4h-2v7l-2.5-1.5L13 11V4H8v12Zm-2 2V2h16v16H6Zm-4 4V6h2v14h14v2H2ZM8 4v12V4Zm5 7l2.5-1.5L18 11l-2.5-1.5L13 11Z"/>`),
		g.Group(children),
	)
}