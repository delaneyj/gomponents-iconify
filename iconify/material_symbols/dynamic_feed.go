package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DynamicFeed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19v-7h2v7h9v2H4Zm4-4q-.825 0-1.413-.588T6 15V8h2v7h9v2H8Zm4-4q-.825 0-1.413-.588T10 11V5q0-.825.588-1.413T12 3h8q.825 0 1.413.588T22 5v6q0 .825-.588 1.413T20 13h-8Zm0-2h8V7h-8v4Z"/>`),
		g.Group(children),
	)
}