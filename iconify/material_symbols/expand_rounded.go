package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.425 0-.713-.288T4 21q0-.425.288-.713T5 20h14q.425 0 .713.288T20 21q0 .425-.288.713T19 22H5Zm7-3.425q-.2 0-.375-.063T11.3 18.3l-2.6-2.6q-.275-.275-.287-.687T8.7 14.3q.275-.275.688-.288t.712.263l.9.875v-6.3l-.9.875Q9.825 10 9.412 10T8.7 9.7q-.275-.275-.275-.7t.275-.7l2.6-2.6q.15-.15.325-.212T12 5.425q.2 0 .375.063t.325.212l2.6 2.6q.275.275.287.688T15.3 9.7q-.275.275-.688.288t-.712-.263L13 8.85v6.3l.9-.875q.275-.275.688-.275t.712.3q.275.275.275.7t-.275.7l-2.6 2.6q-.15.15-.325.213t-.375.062ZM5 4q-.425 0-.713-.288T4 3q0-.425.288-.713T5 2h14q.425 0 .713.288T20 3q0 .425-.288.713T19 4H5Z"/>`),
		g.Group(children),
	)
}