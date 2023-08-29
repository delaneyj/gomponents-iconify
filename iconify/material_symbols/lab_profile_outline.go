package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabProfileOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 12v-2h8v2H8Zm0-4V6h8v2H8Zm-2 6h7.5q.725 0 1.35.313t1.05.887l2.1 2.75V4H6v10Zm0 6h11.05l-2.725-3.575q-.15-.2-.362-.313T13.5 16H6v4Zm12 2H6q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22ZM6 20V4v16Zm0-4v-2v2Z"/>`),
		g.Group(children),
	)
}