package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastFoodEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M10 8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2h9zm0-3H1a1 1 0 1 0 0 2h9a1 1 0 1 0 0-2zM8.55 1H2.46A1.46 1.46 0 0 0 1 2.46V4h9V2.47A1.46 1.46 0 0 0 8.55 1zm-5 2A.5.5 0 1 1 4 2.5a.5.5 0 0 1-.5.5h.05zm4 0A.5.5 0 1 1 8 2.5a.5.5 0 0 1-.5.5h.05z" fill="currentColor"/>`),
		g.Group(children),
	)
}