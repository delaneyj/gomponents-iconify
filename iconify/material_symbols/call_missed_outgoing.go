package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallMissedOutgoing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 17.425l-9-9L4.4 7l7.6 7.6L17.6 9H13V7h8v8h-2v-4.575l-7 7Z"/>`),
		g.Group(children),
	)
}