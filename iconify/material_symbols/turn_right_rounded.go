package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnRightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 19v-8q0-.825.588-1.413T9 9h8.2l-.9-.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l.9-.9H9v8q0 .425-.288.713T8 20q-.425 0-.713-.288T7 19Z"/>`),
		g.Group(children),
	)
}