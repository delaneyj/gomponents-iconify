package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyChangingStationRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.425 0-.713-.288T3 21v-8.675q0-.175.025-.337t.075-.313l1.475-4.3q.275-.85 1.1-1.2t1.625 0L11.45 8H13q.425 0 .713.288T14 9q0 .425-.288.713T13 10h-1.575q-.2 0-.4-.05t-.4-.125l-2.325-1L7 12.75V21q0 .425-.287.713T6 22H4ZM8 5q-.825 0-1.413-.588T6 3q0-.825.588-1.413T8 1q.825 0 1.413.588T10 3q0 .825-.588 1.413T8 5Zm2 14q-.425 0-.713-.288T9 18q0-.425.288-.713T10 17h10q.425 0 .713.288T21 18q0 .425-.288.713T20 19H10Zm9.5-3q-.625 0-1.063-.438T18 14.5q0-.625.438-1.063T19.5 13q.625 0 1.063.438T21 14.5q0 .625-.438 1.063T19.5 16ZM13 16q-.825 0-1.413-.588T11 14v-1h-1q-.425 0-.713-.288T9 12q0-.425.288-.713T10 11h2q.425 0 .713.288T13 12v1h2v-1q0-.425.288-.713T16 11q.425 0 .713.288T17 12v2q0 .825-.588 1.413T15 16h-2Z"/>`),
		g.Group(children),
	)
}