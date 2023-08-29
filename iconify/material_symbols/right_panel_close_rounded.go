package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightPanelCloseRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 14.8q0 .35.3.475t.55-.125l2.45-2.45q.3-.3.3-.7t-.3-.7L8.35 8.85q-.25-.25-.55-.125t-.3.475v5.6ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm9-2V5H5v14h9Z"/>`),
		g.Group(children),
	)
}