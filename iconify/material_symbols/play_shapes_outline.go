package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayShapesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.4 10.5L13 9.1l2.6-2.6L13 3.9l1.4-1.4L17 5.1l2.6-2.6L21 3.9l-2.6 2.6L21 9.1l-1.4 1.4L17 7.9l-2.6 2.6ZM2 11l5-9l5 9H2Zm5 10q-1.65 0-2.825-1.175T3 17q0-1.675 1.175-2.838T7 13q1.65 0 2.825 1.175T11 17q0 1.65-1.175 2.825T7 21Zm0-2q.825 0 1.413-.588T9 17q0-.825-.588-1.413T7 15q-.825 0-1.413.588T5 17q0 .825.588 1.413T7 19ZM5.4 9h3.2L7 6.125L5.4 9ZM13 21v-8h8v8h-8Zm2-2h4v-4h-4v4ZM7 7.55ZM7 17Zm10 0Z"/>`),
		g.Group(children),
	)
}