package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChurchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-4.7q0-.6.325-1.088t.875-.737l2.8-1.25v-2q0-.575.3-1.038t.8-.737L11 6.5V5h-1q-.425 0-.713-.288T9 4q0-.425.288-.713T10 3h1V2q0-.425.288-.713T12 1q.425 0 .713.288T13 2v1h1q.425 0 .713.288T15 4q0 .425-.288.713T14 5h-1v1.5l3.9 1.95q.5.275.8.738t.3 1.037v2l2.8 1.25q.55.25.875.738T22 15.3V20q0 .825-.588 1.413T20 22h-5q-.425 0-.713-.288T14 21v-2q0-.825-.588-1.413T12 17q-.825 0-1.413.588T10 19v2q0 .425-.288.713T9 22H4q-.825 0-1.413-.588T2 20Zm10-6.5q.625 0 1.063-.438T13.5 12q0-.625-.438-1.063T12 10.5q-.625 0-1.063.438T10.5 12q0 .625.438 1.063T12 13.5Z"/>`),
		g.Group(children),
	)
}