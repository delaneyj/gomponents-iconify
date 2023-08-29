package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutpatientMedOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5q-.425 0-.713-.288T2 4q0-.425.288-.713T3 3h10q.425 0 .713.288T14 4q0 .425-.288.713T13 5H3Zm5 12.5q.625 0 1.063-.438T9.5 16v-1h1q.625 0 1.063-.438T12 13.5q0-.625-.438-1.063T10.5 12h-1v-1q0-.625-.438-1.063T8 9.5q-.625 0-1.063.438T6.5 11v1h-1q-.625 0-1.063.438T4 13.5q0 .625.438 1.063T5.5 15h1v1q0 .625.438 1.063T8 17.5ZM3 21q-.825 0-1.413-.588T1 19V8q0-.825.588-1.413T3 6h10q.825 0 1.413.588T15 8v11q0 .825-.588 1.413T13 21H3Zm0-2h10V8H3v11Zm15.775-4.2q-.3-.3-.3-.713t.3-.712L19.15 13H17q-.425 0-.713-.288T16 12q0-.425.288-.713T17 11h2.15l-.375-.375q-.3-.3-.288-.713T18.8 9.2q.3-.3.713-.3t.712.3l2.075 2.1q.3.3.3.7t-.3.7l-2.1 2.1q-.3.3-.713.3t-.712-.3ZM8 13.5Z"/>`),
		g.Group(children),
	)
}