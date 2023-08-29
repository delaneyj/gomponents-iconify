package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangularFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="m67 24l-31 9.5L5 43V5l31 9.5L67 24z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m67 24l-31 9.5L5 43V5l31 9.5L67 24zM5 5v62"/>`),
		g.Group(children),
	)
}