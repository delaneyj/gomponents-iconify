package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommercialEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M0 5a4.71 4.71 0 0 1 1-3h9a4.71 4.71 0 0 1 1 3H0zm2 1v5h4V7h2v4h1V6H2zm3 3H3V7h2v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}