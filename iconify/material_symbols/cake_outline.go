package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.425 0-.713-.288T3 21v-5q0-.825.588-1.413T5 14v-4q0-.825.588-1.413T7 8h4V6.55q-.45-.3-.725-.725T10 4.8q0-.375.15-.738t.45-.662L12 2l1.4 1.4q.3.3.45.662T14 4.8q0 .6-.275 1.025T13 6.55V8h4q.825 0 1.413.587T19 10v4q.825 0 1.413.588T21 16v5q0 .425-.288.713T20 22H4Zm3-8h10v-4H7v4Zm-2 6h14v-4H5v4Zm2-6h10H7Zm-2 6h14H5Zm14-6H5h14Z"/>`),
		g.Group(children),
	)
}