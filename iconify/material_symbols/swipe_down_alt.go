package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeDownAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 22l-4-4l1.4-1.425l1.6 1.6V13.9q-1.725-.35-2.863-1.713T7 9q0-2.075 1.463-3.538T12 4q2.075 0 3.538 1.463T17 9q0 1.825-1.137 3.188T13 13.9v4.275l1.6-1.575L16 18l-4 4Z"/>`),
		g.Group(children),
	)
}