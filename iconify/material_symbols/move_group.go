package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18q-.825 0-1.413-.588T6 16v-2h2v2h12V6H8v2H6V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm-4 4q-.825 0-1.413-.588T2 20V6h2v14h14v2H4Zm9-7l-1.4-1.4l1.575-1.6H6v-2h7.175L11.6 8.4L13 7l4 4l-4 4Z"/>`),
		g.Group(children),
	)
}