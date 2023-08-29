package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InpatientOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h9q.825 0 1.413.588T15 4v16q0 .825-.588 1.413T13 22H4Zm0-11.475q.45-.275.95-.4T6 10h5q.55 0 1.05.125t.95.4V4H4v6.525ZM8.5 9q-.825 0-1.413-.588T6.5 7q0-.825.588-1.413T8.5 5q.825 0 1.413.588T10.5 7q0 .825-.588 1.413T8.5 9Zm11 6.5L16 12l3.5-3.5l1.4 1.4l-1.075 1.1H23v2h-3.175l1.075 1.1l-1.4 1.4ZM4 20h9v-6q0-.825-.588-1.413T11 12H6q-.825 0-1.413.588T4 14v6Zm3.5-1h2v-2h2v-2h-2v-2h-2v2h-2v2h2v2ZM4 20h9h-9Z"/>`),
		g.Group(children),
	)
}