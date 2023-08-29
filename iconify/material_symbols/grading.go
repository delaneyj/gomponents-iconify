package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grading(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.425 20.975L13.6 18.15l1.4-1.4l1.425 1.425L19.6 15l1.4 1.4l-4.575 4.575ZM3 21v-2h9v2H3Zm0-4v-2h9v2H3Zm0-4v-2h18v2H3Zm0-4V7h18v2H3Zm0-4V3h18v2H3Z"/>`),
		g.Group(children),
	)
}