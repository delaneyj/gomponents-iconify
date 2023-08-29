package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReopenWindowRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v6q0 .425-.288.713T21 13q-.425 0-.713-.288T20 12V8H4v10h7q.425 0 .713.288T12 19q0 .425-.288.713T11 20H4Zm15 4q-1.425 0-2.6-.713t-1.825-1.912q-.175-.3-.062-.638t.437-.462q.275-.125.55 0t.425.4q.45.825 1.275 1.325t1.8.5q1.45 0 2.475-1.025T22.5 19q0-1.45-1.025-2.475T19 15.5q-.7 0-1.325.25t-1.125.75h.7q.325 0 .538.213t.212.537q0 .325-.213.537T17.25 18H15q-.425 0-.713-.288T14 17v-2.25q0-.325.213-.537T14.75 14q.325 0 .537.213t.213.537v.675q.725-.675 1.625-1.05T19 14q2.075 0 3.538 1.463T24 19q0 2.075-1.463 3.538T19 24Z"/>`),
		g.Group(children),
	)
}