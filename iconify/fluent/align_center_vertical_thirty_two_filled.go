package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVerticalThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 30a1 1 0 0 1-1-1v-2h-3.5A3.5 3.5 0 0 1 8 23.5v-3a3.5 3.5 0 0 1 3.5-3.5H15v-2H8.5A3.5 3.5 0 0 1 5 11.5v-3A3.5 3.5 0 0 1 8.5 5H15V3a1 1 0 1 1 2 0v2h6.5A3.5 3.5 0 0 1 27 8.5v3a3.5 3.5 0 0 1-3.5 3.5H17v2h3.5a3.5 3.5 0 0 1 3.5 3.5v3a3.5 3.5 0 0 1-3.5 3.5H17v2a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}