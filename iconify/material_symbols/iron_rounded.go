package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IronRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18q-.425 0-.713-.288T2 17v-2q0-1.65 1.175-2.825T6 11h9v-1q0-.425-.288-.713T14 9h-4q-.275 0-.5.125t-.35.35q-.15.25-.375.388t-.5.137q-.575 0-.863-.488T7.35 8.6q.375-.725 1.075-1.163T10 7h4q1.25 0 2.125.875T17 10v4q.425 0 .713-.288T18 13V9q0-1.25.875-2.125T21 6q.425 0 .713.288T22 7q0 .425-.288.713T21 8t-.713.288Q20 8.575 20 9v4q0 1.25-.875 2.125T17 16v1q0 .425-.288.713T16 18H3Z"/>`),
		g.Group(children),
	)
}