package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicationLiquidOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5q-.425 0-.713-.288T3 4q0-.425.288-.713T4 3h10q.425 0 .713.288T15 4q0 .425-.288.713T14 5H4Zm5 12.5q.625 0 1.063-.438T10.5 16v-1h1q.625 0 1.063-.438T13 13.5q0-.625-.438-1.063T11.5 12h-1v-1q0-.625-.438-1.063T9 9.5q-.625 0-1.063.438T7.5 11v1h-1q-.625 0-1.063.438T5 13.5q0 .625.438 1.063T6.5 15h1v1q0 .625.438 1.063T9 17.5ZM4 21q-.825 0-1.413-.588T2 19V8q0-.825.588-1.413T4 6h10q.825 0 1.413.588T16 8v11q0 .825-.588 1.413T14 21H4Zm0-2h10V8H4v11Zm16 2q-.425 0-.713-.288T19 20v-6.25q-.875-.425-1.438-1.413T17 10q0-1.7.863-2.85T20 6q1.275 0 2.138 1.15T23 10q0 1.35-.563 2.337T21 13.75V20q0 .425-.288.713T20 21Zm0-9q.3 0 .65-.537T21 10q0-.925-.35-1.463T20 8q-.3 0-.65.537T19 10q0 .925.35 1.463T20 12ZM9 13.5ZM20 10Z"/>`),
		g.Group(children),
	)
}