package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitBranchSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 2 4.95v5.1A2.5 2.5 0 1 0 4.95 13H9.5A3.5 3.5 0 0 0 13 9.5V7.95a2.5 2.5 0 1 0-1 0V9.5A2.5 2.5 0 0 1 9.5 12H4.95A2.503 2.503 0 0 0 3 10.05v-5.1A2.5 2.5 0 0 0 2.5 0Z"/>`),
		g.Group(children),
	)
}