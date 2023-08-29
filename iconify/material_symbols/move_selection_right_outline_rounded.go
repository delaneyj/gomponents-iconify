package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveSelectionRightOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 8q0-.825.588-1.413T12 6h8q.825 0 1.413.588T22 8v8q0 .825-.588 1.413T20 18h-8q-.825 0-1.413-.588T10 16V8Zm2 0v8h8V8h-8ZM6 18v-2h2v2H6ZM6 8V6h2v2H6Zm-4 5v-2h2v2H2Zm14-1ZM2 16h2v2q-.825 0-1.413-.588T2 16Zm0-8q0-.825.588-1.413T4 6v2H2Z"/>`),
		g.Group(children),
	)
}