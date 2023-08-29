package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19v-9q0-.425.288-.713T2 9q.425 0 .713.288T3 10v1h2V4q0-.425.288-.713T6 3q.425 0 .713.288T7 4v1h2V4q0-.425.288-.713T10 3q.425 0 .713.288T11 4v1h2V4q0-.425.288-.713T14 3q.425 0 .713.288T15 4v1h2V4q0-.425.288-.713T18 3q.425 0 .713.288T19 4v7h2v-1q0-.425.288-.713T22 9q.425 0 .713.288T23 10v9q0 .825-.588 1.413T21 21h-6q-.425 0-.713-.288T14 20v-2q0-.825-.588-1.413T12 16q-.825 0-1.413.588T10 18v2q0 .425-.288.713T9 21H3q-.825 0-1.413-.588T1 19Zm2 0h5v-1q0-1.65 1.175-2.825T12 14q1.65 0 2.825 1.175T16 18v1h5v-6h-3q-.425 0-.713-.288T17 12V7H7v5q0 .425-.288.713T6 13H3v6Zm6.25-7h1.5q.125 0 .188-.063T11 11.75V10q0-.425-.288-.713T10 9q-.425 0-.713.288T9 10v1.75q0 .125.063.188T9.25 12Zm4 0h1.5q.125 0 .188-.063T15 11.75V10q0-.425-.288-.713T14 9q-.425 0-.713.288T13 10v1.75q0 .125.063.188t.187.062ZM12 13Z"/>`),
		g.Group(children),
	)
}