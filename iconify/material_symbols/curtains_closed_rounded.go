package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurtainsClosedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19V5q0-.825.588-1.413T6 3h12q.825 0 1.413.588T20 5v14h1q.425 0 .713.288T22 20q0 .425-.288.713T21 21H3q-.425 0-.713-.288T2 20q0-.425.288-.713T3 19h1Zm7 0h2V5h-2v14Z"/>`),
		g.Group(children),
	)
}