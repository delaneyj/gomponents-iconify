package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PunchClockOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20V8q0-.825.588-1.413T5 6h1V3q0-.825.588-1.413T8 1h8q.825 0 1.413.588T18 3v3h1q.825 0 1.413.588T21 8v12q0 .825-.588 1.413T19 22H5ZM8 6h8V3H8v3ZM5 20h14V8H5v12Zm7-1q-2.075 0-3.538-1.463T7 14q0-2.075 1.463-3.538T12 9q2.075 0 3.538 1.463T17 14q0 2.075-1.463 3.538T12 19Zm0-1.5q1.45 0 2.475-1.025T15.5 14q0-1.45-1.025-2.475T12 10.5q-1.45 0-2.475 1.025T8.5 14q0 1.45 1.025 2.475T12 17.5Zm.8-2l-1.15-1.15q-.05-.05-.15-.35v-2q0-.2.15-.35t.35-.15q.2 0 .35.15t.15.35v1.8l1 1q.15.15.15.35t-.15.35q-.15.15-.35.15t-.35-.15ZM12 14Z"/>`),
		g.Group(children),
	)
}