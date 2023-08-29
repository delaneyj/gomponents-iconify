package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m1.15 6.69l7.71-5.14l.7 1l-7.12 4.78H15.5v1.25H2.74l6.83 4.86l-.72 1L1.15 9A1.42 1.42 0 0 1 .5 7.83a1.4 1.4 0 0 1 .65-1.14z"/>`),
		g.Group(children),
	)
}