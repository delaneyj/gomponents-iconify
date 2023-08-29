package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.106 12.69a1.49 1.49 0 0 1-2.104 0L1.561 6.246c-.582-.581-.839-2.103 1-2.103h12.988c1.901 0 1.582 1.521 1 2.103l-6.443 6.444z"/>`),
		g.Group(children),
	)
}