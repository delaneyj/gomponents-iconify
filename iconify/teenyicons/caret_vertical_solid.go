package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretVerticalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 1.336L2.17 6h10.66L7.5 1.336Zm0 12.328L12.83 9H2.17l5.33 4.664Z"/>`),
		g.Group(children),
	)
}