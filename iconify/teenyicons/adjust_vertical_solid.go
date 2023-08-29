package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustVerticalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2 0v6H0v3h2v6h1V9h2V6H3V0H2Zm3 10h2V0h1v10h2v3H8v2H7v-2H5v-3Zm7-10v2h-2v3h2v10h1V5h2V2h-2V0h-1Z"/>`),
		g.Group(children),
	)
}