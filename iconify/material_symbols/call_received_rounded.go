package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallReceivedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 20q-.425 0-.713-.288T5 19v-8q0-.425.288-.713T6 10q.425 0 .713.288T7 11v5.6L17.925 5.675Q18.2 5.4 18.6 5.4t.7.3q.275.275.275.7t-.275.7L8.4 18H14q.425 0 .713.288T15 19q0 .425-.288.713T14 20H6Z"/>`),
		g.Group(children),
	)
}