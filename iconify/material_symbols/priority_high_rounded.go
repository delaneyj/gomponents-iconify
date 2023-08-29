package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PriorityHighRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-.825 0-1.413-.588T10 19q0-.825.588-1.413T12 17q.825 0 1.413.588T14 19q0 .825-.588 1.413T12 21Zm0-6q-.825 0-1.413-.588T10 13V5q0-.825.588-1.413T12 3q.825 0 1.413.588T14 5v8q0 .825-.588 1.413T12 15Z"/>`),
		g.Group(children),
	)
}