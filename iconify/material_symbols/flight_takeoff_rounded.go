package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightTakeoffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 21H4q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h16q.425 0 .713.288T21 20q0 .425-.288.713T20 21ZM5.3 12.4l4.8-1.3l-3.475-5.9q-.2-.35-.075-.75t.55-.525l.425-.125q.225-.075.45-.025t.4.2L14.85 9.8l5-1.35q.8-.225 1.45.3t.65 1.4q0 .55-.338.975t-.862.575L5.575 15.8q-.325.1-.625-.025t-.475-.425L2.45 11.9q-.175-.275-.038-.575t.463-.35l.375-.075q.15-.025.275.013t.25.137L5.3 12.4Z"/>`),
		g.Group(children),
	)
}