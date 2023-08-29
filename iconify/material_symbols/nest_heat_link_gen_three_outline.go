package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestHeatLinkGenThreeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21q-2.075 0-3.538-1.463T3 16V8q0-2.075 1.463-3.538T8 3h8q2.075 0 3.538 1.463T21 8v8q0 2.075-1.463 3.538T16 21H8Zm0-2h8q1.25 0 2.125-.875T19 16V8q0-1.25-.875-2.125T16 5H8q-1.25 0-2.125.875T5 8v8q0 1.25.875 2.125T8 19Zm4-2q1.875 0 3.188-1.313T16.5 12.5q0-1.875-1.313-3.188T12 8q-1.875 0-3.188 1.313T7.5 12.5q0 1.875 1.313 3.188T12 17Zm0-2q-1.05 0-1.775-.725T9.5 12.5q0-1.05.725-1.775T12 10q1.05 0 1.775.725T14.5 12.5q0 1.05-.725 1.775T12 15Zm0-7.5q.425 0 .713-.288T13 6.5q0-.425-.288-.713T12 5.5q-.425 0-.713.288T11 6.5q0 .425.288.713T12 7.5Zm0 5Z"/>`),
		g.Group(children),
	)
}