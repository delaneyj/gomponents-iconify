package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorbellThreeP(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22q-.825 0-1.413-.588T5 20V4q0-.825.588-1.413T7 2h10q.825 0 1.413.588T19 4v16q0 .825-.588 1.413T17 22H7Zm5-8.5q.45 0 .725-.275T13 12.5h-2q0 .45.275.725T12 13.5ZM8 12h8v-1h-1V8.7q0-1.125-.575-2.013T12.8 5.5v-.3q0-.35-.225-.575T12 4.4q-.35 0-.575.225T11.2 5.2v.3q-1.05.375-1.625 1.225T9 8.7V11H8v1Zm4 7q.825 0 1.413-.588T14 17q0-.825-.588-1.413T12 15q-.825 0-1.413.588T10 17q0 .825.588 1.413T12 19Zm0-1q-.425 0-.713-.288T11 17q0-.425.288-.713T12 16q.425 0 .713.288T13 17q0 .425-.288.713T12 18Z"/>`),
		g.Group(children),
	)
}