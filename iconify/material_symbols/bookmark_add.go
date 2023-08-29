package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21V5q0-.825.588-1.413T7 3h7q-.5.75-.75 1.438T13 6q0 1.8 1.137 3.175T17 10.9q.575.075 1 .075t1-.075V21l-7-3l-7 3ZM17 9V7h-2V5h2V3h2v2h2v2h-2v2h-2Z"/>`),
		g.Group(children),
	)
}