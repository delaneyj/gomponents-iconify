package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GMobiledata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17q-.825 0-1.413-.588T7 15V9q0-.825.588-1.413T9 7h5q.825 0 1.413.588T16 9H9v6h5v-2h-2v-2h4v4q0 .825-.588 1.413T14 17H9Z"/>`),
		g.Group(children),
	)
}