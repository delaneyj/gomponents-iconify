package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseSidingRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19v-7H3.3q-.35 0-.475-.325t.15-.55l8.35-7.525q.275-.275.675-.275t.675.275l8.35 7.525q.275.225.15.55T20.7 12H19v7q0 .425-.288.713T18 20q-.425 0-.713-.288T17 19v-1H7v1q0 .425-.288.713T6 20q-.425 0-.713-.288T5 19ZM9.45 8h5.1L12 5.7L9.45 8ZM7 12h10v-1.8l-.225-.2h-9.55L7 10.2V12Zm0 4h10v-2H7v2Z"/>`),
		g.Group(children),
	)
}