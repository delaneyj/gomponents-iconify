package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScubaDiving(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 23l-1.6-1.2L5.25 18l.775-4.45q.075-.6.475-1.063t1.025-.612L17 9l2-4l3-3l1 1l-2.5 2.9l-1.5 4.6l-5 3.5l-5.85 1.85L7 19l-3 4Zm-1-8q-.825 0-1.413-.588T1 13q0-.825.588-1.413T3 11q.825 0 1.413.588T5 13q0 .825-.588 1.413T3 15Zm5.9-4.9q-.6.175-1.137-.138T7.05 9.05q-.175-.6.138-1.15t.912-.7L12.65 6l.775 2.9L8.9 10.1Z"/>`),
		g.Group(children),
	)
}