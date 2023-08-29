package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkAddedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.825 9L15 6.175l1.4-1.425l1.425 1.425l3.525-3.55l1.425 1.425L17.825 9ZM5 21V5q0-.825.588-1.413T7 3h6v2H7v12.95l5-2.15l5 2.15V11h2v10l-7-3l-7 3ZM7 5h6h-6Z"/>`),
		g.Group(children),
	)
}