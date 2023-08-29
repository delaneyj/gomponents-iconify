package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 213V0h171v213H0zm0 171V256h171v128H0zm213 0V171h171v213H213zm0-384h171v128H213V0z"/>`),
		g.Group(children),
	)
}