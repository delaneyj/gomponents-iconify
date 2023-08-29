package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsLevelOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17q-.825 0-1.413-.588T2 15V9q0-.825.588-1.413T4 7h16q.825 0 1.413.588T22 9v6q0 .825-.588 1.413T20 17H4Zm0-2h16V9h-2.35q.175.35.263.725T18 10.5q0 1.425-1.038 2.463T14.5 14h-5q-1.425 0-2.463-1.038T6 10.5q0-.4.088-.775T6.35 9H4v6Zm5.5-3H11V9H9.5q-.6 0-1.05.45T8 10.5q0 .6.45 1.05T9.5 12Zm3.5 0h1.5q.6 0 1.05-.45T16 10.5q0-.6-.45-1.05T14.5 9H13v3Zm7 3H4h16Z"/>`),
		g.Group(children),
	)
}