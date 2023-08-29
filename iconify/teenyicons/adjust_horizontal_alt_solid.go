package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustHorizontalAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4.5 1a2.5 2.5 0 0 0-2.45 2H0v1h2.05a2.5 2.5 0 0 0 4.9 0H15V3H6.95A2.5 2.5 0 0 0 4.5 1Zm6 8a2.5 2.5 0 0 0-2.45 2H0v1h8.05a2.5 2.5 0 0 0 4.9 0H15v-1h-2.05a2.5 2.5 0 0 0-2.45-2Z"/>`),
		g.Group(children),
	)
}