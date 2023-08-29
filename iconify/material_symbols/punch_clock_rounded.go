package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PunchClockRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20V8q0-.825.588-1.413T5 6h1V3q0-.825.588-1.413T8 1h8q.825 0 1.413.588T18 3v3h1q.825 0 1.413.588T21 8v12q0 .825-.588 1.413T19 22H5ZM8 6h8V3H8v3Zm4 12.2q1.725 0 2.963-1.237T16.2 14q0-1.725-1.238-2.963T12 9.8q-1.725 0-2.963 1.238T7.8 14q0 1.725 1.238 2.963T12 18.2Zm.8-2.7l-1.15-1.15q-.05-.05-.15-.35v-2q0-.2.15-.35t.35-.15q.2 0 .35.15t.15.35v1.8l1 1q.15.15.15.35t-.15.35q-.15.15-.35.15t-.35-.15Z"/>`),
		g.Group(children),
	)
}