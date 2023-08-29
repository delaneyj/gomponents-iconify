package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkAddSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 9V7h-2V5h2V3h2v2h2v2h-2v2h-2ZM5 21V3h9q-.5.75-.75 1.438T13 6q0 1.8 1.137 3.175T17 10.9q.575.075 1 .075t1-.075V21l-7-3l-7 3Z"/>`),
		g.Group(children),
	)
}