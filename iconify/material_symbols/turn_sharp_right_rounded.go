package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnSharpRightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21q-.425 0-.713-.288T6 20v-5q0-.825.588-1.413T8 13h8V6.8l-.9.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l2.6-2.6q.3-.3.7-.3t.7.3l2.6 2.6q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-.9-.9V13q0 .825-.587 1.413T16 15H8v5q0 .425-.288.713T7 21Z"/>`),
		g.Group(children),
	)
}