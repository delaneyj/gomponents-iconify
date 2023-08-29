package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MilitaryTechRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.15 21.05q-.275.225-.588.013t-.187-.563l.725-2.3l-1.825-1.3q-.3-.2-.175-.55t.475-.35H10.8l.7-2.3l-3.55-2.1q-.45-.275-.7-.725T7 9.85V4q0-.825.588-1.413T9 2h6q.825 0 1.413.588T17 4v5.85q0 .575-.25 1.025t-.7.725l-3.55 2.1l.7 2.3h2.225q.35 0 .475.35t-.175.55L13.9 18.2l.725 2.3q.125.35-.188.563t-.587-.013L12 19.65l-1.85 1.4Zm1.85-9.4l1-.6V5q0-.425-.288-.713T12 4q-.425 0-.713.288T11 5v6.05l1 .6Z"/>`),
		g.Group(children),
	)
}