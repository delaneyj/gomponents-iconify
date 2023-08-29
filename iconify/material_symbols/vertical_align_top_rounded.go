package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalAlignTopRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5q-.425 0-.713-.288T4 4q0-.425.288-.713T5 3h14q.425 0 .713.288T20 4q0 .425-.288.713T19 5H5Zm7 16q-.425 0-.713-.288T11 20v-9.2l-1.9 1.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l3.6-3.6q.15-.15.325-.212T12 7.425q.2 0 .375.063t.325.212l3.6 3.6q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L13 10.8V20q0 .425-.287.713T12 21Z"/>`),
		g.Group(children),
	)
}