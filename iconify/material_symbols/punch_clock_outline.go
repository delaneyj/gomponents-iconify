package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PunchClockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20V8q0-.825.588-1.413T5 6h1V1h12v5h1q.825 0 1.413.588T21 8v12q0 .825-.588 1.413T19 22H5ZM8 6h8V3H8v3ZM5 20h14V8H5v12Zm7-1q-2.075 0-3.538-1.463T7 14q0-2.075 1.463-3.538T12 9q2.075 0 3.538 1.463T17 14q0 2.075-1.463 3.538T12 19Zm0-1.5q1.45 0 2.475-1.025T15.5 14q0-1.45-1.025-2.475T12 10.5q-1.45 0-2.475 1.025T8.5 14q0 1.45 1.025 2.475T12 17.5Zm1.15-1.65L11.5 14.2v-2.7h1v2.3l1.35 1.35l-.7.7ZM12 14Z"/>`),
		g.Group(children),
	)
}