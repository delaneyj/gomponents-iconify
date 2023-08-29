package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalAlignCenterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 17.8l-.9.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l2.6-2.6q.15-.15.325-.212t.375-.063q.2 0 .375.063t.325.212l2.6 2.6q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-.9-.9V21q0 .425-.288.713T12 22q-.425 0-.713-.288T11 21v-3.2ZM5 13q-.425 0-.713-.288T4 12q0-.425.288-.713T5 11h14q.425 0 .713.288T20 12q0 .425-.288.713T19 13H5Zm6-6.8V3q0-.425.288-.713T12 2q.425 0 .713.288T13 3v3.2l.9-.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-2.6 2.6q-.15.15-.325.213T12 9.575q-.2 0-.375-.062T11.3 9.3L8.7 6.7q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l.9.9Z"/>`),
		g.Group(children),
	)
}