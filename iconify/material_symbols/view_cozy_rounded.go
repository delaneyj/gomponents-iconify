package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCozyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 10.5H5q-.825 0-1.413-.588T3 8.5V5q0-.825.588-1.413T5 3h3.5q.825 0 1.413.588T10.5 5v3.5q0 .825-.588 1.413T8.5 10.5Zm0 10.5H5q-.825 0-1.413-.588T3 19v-3.5q0-.825.588-1.413T5 13.5h3.5q.825 0 1.413.588T10.5 15.5V19q0 .825-.588 1.413T8.5 21ZM19 10.5h-3.5q-.825 0-1.413-.588T13.5 8.5V5q0-.825.588-1.413T15.5 3H19q.825 0 1.413.588T21 5v3.5q0 .825-.588 1.413T19 10.5ZM19 21h-3.5q-.825 0-1.413-.588T13.5 19v-3.5q0-.825.588-1.413T15.5 13.5H19q.825 0 1.413.588T21 15.5V19q0 .825-.588 1.413T19 21Z"/>`),
		g.Group(children),
	)
}