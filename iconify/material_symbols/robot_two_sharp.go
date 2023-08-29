package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotTwoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-7h16v7H4Zm5-8q-2.075 0-3.538-1.463T4 8q0-2.075 1.463-3.538T9 3h6q2.075 0 3.538 1.463T20 8q0 2.075-1.463 3.538T15 13H9Zm0-4q.425 0 .713-.288T10 8q0-.425-.288-.713T9 7q-.425 0-.713.288T8 8q0 .425.288.713T9 9Zm6 0q.425 0 .713-.288T16 8q0-.425-.288-.713T15 7q-.425 0-.713.288T14 8q0 .425.288.713T15 9Z"/>`),
		g.Group(children),
	)
}