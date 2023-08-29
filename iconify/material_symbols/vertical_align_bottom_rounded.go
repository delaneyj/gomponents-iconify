package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalAlignBottomRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.425 0-.713-.288T4 20q0-.425.288-.713T5 19h14q.425 0 .713.288T20 20q0 .425-.288.713T19 21H5Zm7-4.425q-.2 0-.375-.063T11.3 16.3l-3.6-3.6q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.9 1.9V4q0-.425.288-.713T12 3q.425 0 .713.288T13 4v9.2l1.9-1.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-3.6 3.6q-.15.15-.325.213t-.375.062Z"/>`),
		g.Group(children),
	)
}