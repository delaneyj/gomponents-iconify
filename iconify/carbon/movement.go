package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Movement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24 20l-1.41 1.41L26.17 25H10a4 4 0 0 1 0-8h12a6 6 0 0 0 0-12H5.83l3.58-3.59L8 0L2 6l6 6l1.41-1.41L5.83 7H22a4 4 0 0 1 0 8H10a6 6 0 0 0 0 12h16.17l-3.58 3.59L24 32l6-6Z"/>`),
		g.Group(children),
	)
}