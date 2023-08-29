package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideogameAssetRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18q-.825 0-1.413-.588T2 16V8q0-.825.588-1.413T4 6h16q.825 0 1.413.588T22 8v8q0 .825-.588 1.413T20 18H4Zm3-5v1q0 .425.288.713T8 15q.425 0 .713-.288T9 14v-1h1q.425 0 .713-.288T11 12q0-.425-.288-.713T10 11H9v-1q0-.425-.288-.713T8 9q-.425 0-.713.288T7 10v1H6q-.425 0-.713.288T5 12q0 .425.288.713T6 13h1Zm7.5 2q.625 0 1.063-.438T16 13.5q0-.625-.438-1.063T14.5 12q-.625 0-1.063.438T13 13.5q0 .625.438 1.063T14.5 15Zm3-3q.625 0 1.063-.438T19 10.5q0-.625-.438-1.063T17.5 9q-.625 0-1.063.438T16 10.5q0 .625.438 1.063T17.5 12Z"/>`),
		g.Group(children),
	)
}