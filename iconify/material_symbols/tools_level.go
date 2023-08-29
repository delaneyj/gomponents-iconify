package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsLevel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17q-.825 0-1.413-.588T2 15V9q0-.825.588-1.413T4 7h16q.825 0 1.413.588T22 9v6q0 .825-.588 1.413T20 17H4Zm5.5-5H11V9H9.5q-.6 0-1.05.45T8 10.5q0 .6.45 1.05T9.5 12Zm3.5 0h1.5q.6 0 1.05-.45T16 10.5q0-.6-.45-1.05T14.5 9H13v3Z"/>`),
		g.Group(children),
	)
}