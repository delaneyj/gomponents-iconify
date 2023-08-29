package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.425 0-.713-.288T3 21v-4q0-.825.588-1.413T5 15h14q.825 0 1.413.588T21 17v4q0 .425-.288.713T20 22H4Zm1-9v-3q0-.825.588-1.413T7 8h4V6.55q-.45-.3-.725-.725T10 4.8q0-.375.15-.738t.45-.662L12 2l1.4 1.4q.3.3.45.662T14 4.8q0 .6-.275 1.025T13 6.55V8h4q.825 0 1.413.587T19 10v3H5Z"/>`),
		g.Group(children),
	)
}