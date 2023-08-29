package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicExternalOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 6.65L5.35 2.5q.35-.275.775-.388T7 2q1.25 0 2.125.862T10 5q0 .45-.138.863T9.5 6.65ZM20 17.15l-2-2V6q0-.825-.588-1.413T16 4q-.825 0-1.413.588T14 6v5.15l-2-2V6q0-1.7 1.175-2.85T16 2q1.65 0 2.825 1.15T20 6v11.15Zm.5 6.15L14 16.8V18q0 1.65-1.175 2.825T10 22q-1.65 0-2.825-1.175T6 18H5L4 8h1.15L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM10 20q.825 0 1.413-.588T12 18v-3.2l-2.45-2.45L9 18H8q0 .825.588 1.413T10 20Z"/>`),
		g.Group(children),
	)
}