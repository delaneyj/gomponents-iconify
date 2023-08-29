package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatUnderlinedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-.425 0-.713-.288T5 20q0-.425.288-.713T6 19h12q.425 0 .713.288T19 20q0 .425-.288.713T18 21H6Zm6-4q-2.525 0-3.925-1.575t-1.4-4.175V4.275q0-.525.388-.9T7.975 3q.525 0 .9.375t.375.9V11.4q0 1.4.7 2.275t2.05.875q1.35 0 2.05-.875t.7-2.275V4.275q0-.525.388-.9T16.05 3q.525 0 .9.375t.375.9v6.975q0 2.6-1.4 4.175T12 17Z"/>`),
		g.Group(children),
	)
}