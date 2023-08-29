package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallMissed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 17.425l-7-7V15H3V7h8v2H6.4l5.6 5.6L19.6 7L21 8.425l-9 9Z"/>`),
		g.Group(children),
	)
}