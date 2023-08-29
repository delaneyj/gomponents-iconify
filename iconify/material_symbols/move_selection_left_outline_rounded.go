package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveSelectionLeftOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 16q0 .825-.588 1.413T12 18H4q-.825 0-1.413-.588T2 16V8q0-.825.588-1.413T4 6h8q.825 0 1.413.588T14 8v8Zm-2 0V8H4v8h8Zm4-8V6h2v2h-2Zm0 10v-2h2v2h-2Zm4-5v-2h2v2h-2ZM8 12Zm12-4V6q.825 0 1.413.588T22 8h-2Zm0 8h2q0 .825-.588 1.413T20 18v-2Z"/>`),
		g.Group(children),
	)
}