package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 17l-6-6l6-6l1.425 1.4l-4.6 4.6l4.6 4.6L8 17Zm12 2v-4q0-1.25-.875-2.125T17 12h-6.175l3.6 3.6L13 17l-6-6l6-6l1.425 1.4l-3.6 3.6H17q2.075 0 3.538 1.463T22 15v4h-2Z"/>`),
		g.Group(children),
	)
}