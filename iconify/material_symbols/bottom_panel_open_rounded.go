package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomPanelOpenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.2 11.5h5.6q.35 0 .475-.3t-.125-.55L12.7 8.2q-.3-.3-.7-.3t-.7.3l-2.45 2.45q-.25.25-.125.55t.475.3ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-7h14V5H5v9Z"/>`),
		g.Group(children),
	)
}