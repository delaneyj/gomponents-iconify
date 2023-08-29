package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 1200V0h200v550L700 50v500l500-500v1100L700 650v500L200 650v550H0z"/>`),
		g.Group(children),
	)
}