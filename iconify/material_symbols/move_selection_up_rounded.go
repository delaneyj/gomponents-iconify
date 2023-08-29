package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveSelectionUpRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 14q-.825 0-1.413-.588T6 12V4q0-.825.588-1.413T8 2h8q.825 0 1.413.588T18 4v8q0 .825-.588 1.413T16 14H8Zm8 4v-2h2v2h-2ZM6 18v-2h2v2H6Zm5 4v-2h2v2h-2Zm5 0v-2h2q0 .825-.588 1.413T16 22Zm-8 0q-.825 0-1.413-.588T6 20h2v2Z"/>`),
		g.Group(children),
	)
}