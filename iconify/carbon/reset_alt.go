package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResetAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 8H6.83l3.58-3.59L9 3L3 9l6 6l1.41-1.41L6.83 10H27v16H7v-7H5v7a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}