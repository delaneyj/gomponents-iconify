package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedEmailOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17q-.825 0-1.413-.588T5 15V5q0-.825.588-1.413T7 3h14q.825 0 1.413.588T23 5v10q0 .825-.588 1.413T21 17H7Zm6.425-5.1L7 7.425V15h14V7.425L14.575 11.9q-.275.2-.575.2t-.575-.2ZM14 9.85L21 5H7l7 4.85ZM3 21q-.825 0-1.413-.588T1 19V7.5q0-.425.288-.713T2 6.5q.425 0 .713.288T3 7.5V19h15.5q.425 0 .713.288T19.5 20q0 .425-.288.713T18.5 21H3ZM21 7.35V5H7v2.35V5h14v2.35Z"/>`),
		g.Group(children),
	)
}