package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveSelectionDownRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-.825 0-1.413-.588T6 20v-8q0-.825.588-1.413T8 10h8q.825 0 1.413.588T18 12v8q0 .825-.588 1.413T16 22H8ZM6 8V6h2v2H6Zm10 0V6h2v2h-2Zm-5-4V2h2v2h-2ZM6 4q0-.825.588-1.413T8 2v2H6Zm10 0V2q.825 0 1.413.588T18 4h-2Z"/>`),
		g.Group(children),
	)
}