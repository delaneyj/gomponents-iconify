package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardHideOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 23l-4-4h8l-4 4Zm-8-6q-.825 0-1.413-.588T2 15V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v10q0 .825-.588 1.413T20 17H4Zm0-2h16V5H4v10Zm4-1h8v-2H8v2Zm-3-3h2V9H5v2Zm3 0h2V9H8v2Zm3 0h2V9h-2v2Zm3 0h2V9h-2v2Zm3 0h2V9h-2v2ZM5 8h2V6H5v2Zm3 0h2V6H8v2Zm3 0h2V6h-2v2Zm3 0h2V6h-2v2Zm3 0h2V6h-2v2ZM4 15V5v10Z"/>`),
		g.Group(children),
	)
}