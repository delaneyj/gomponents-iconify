package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 22H7q-.825 0-1.413-.588T5 20V4q0-.825.588-1.413T7 2h10q.825 0 1.413.588T19 4v16q0 .825-.588 1.413T17 22ZM12 9q.825 0 1.413-.588T14 7q0-.825-.588-1.413T12 5q-.825 0-1.413.588T10 7q0 .825.588 1.413T12 9Zm0 10q1.65 0 2.825-1.175T16 15q0-1.65-1.175-2.825T12 11q-1.65 0-2.825 1.175T8 15q0 1.65 1.175 2.825T12 19Zm0-2q-.825 0-1.413-.588T10 15q0-.825.588-1.413T12 13q.825 0 1.413.588T14 15q0 .825-.588 1.413T12 17Z"/>`),
		g.Group(children),
	)
}