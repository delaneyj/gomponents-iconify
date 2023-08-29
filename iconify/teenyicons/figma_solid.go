package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4 9.5a2.496 2.496 0 0 1-1-2c0-.818.393-1.544 1-2A2.5 2.5 0 0 1 5.5 1h4A2.5 2.5 0 0 1 11 5.5a2.5 2.5 0 0 1-3 4v2a2.5 2.5 0 1 1-4-2Z"/>`),
		g.Group(children),
	)
}