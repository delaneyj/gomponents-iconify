package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5.56 14L5 13.587V2.393L5.54 2L11 7.627v.827L5.56 14z"/>`),
		g.Group(children),
	)
}