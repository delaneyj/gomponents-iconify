package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShopSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 1H1V0h13v1ZM1.01 2.402A.5.5 0 0 1 1.5 2h12a.5.5 0 0 1 .49.402l1 5A.5.5 0 0 1 14.5 8H.5a.5.5 0 0 1-.49-.598l1-5ZM1 9v5H0v1h15v-1h-1V9h-2v2H3V9H1Z"/><path fill="currentColor" d="M4 9h7v1H4V9Z"/>`),
		g.Group(children),
	)
}