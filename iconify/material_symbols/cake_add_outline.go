package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeAddOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 8V6h-2V4h2V2h2v2h2v2h-2v2h-2ZM3 22q-.425 0-.713-.288T2 21v-5q0-.825.588-1.413T4 14v-4q0-.825.588-1.413T6 8h4V6.55q-.45-.3-.725-.725T9 4.8q0-.375.15-.738T9.6 3.4L11 2l1.4 1.4q.3.3.45.662T13 4.8q0 .6-.275 1.025T12 6.55V8h4q.825 0 1.413.587T18 10v4q.825 0 1.413.588T20 16v5q0 .425-.288.713T19 22H3Zm3-8h10v-4H6v4Zm-2 6h14v-4H4v4Zm2-6h10H6Zm-2 6h14H4Zm14-6H4h14Z"/>`),
		g.Group(children),
	)
}