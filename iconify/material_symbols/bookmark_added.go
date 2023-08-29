package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkAdded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.825 9L15 6.175l1.4-1.425l1.425 1.425l3.525-3.55l1.425 1.425L17.825 9ZM5 21V5q0-.825.588-1.413T7 3h7q-.5.75-.75 1.438T13 6q0 1.8 1.137 3.175T17 10.9q.575.075 1 .075t1-.075V21l-7-3l-7 3Z"/>`),
		g.Group(children),
	)
}