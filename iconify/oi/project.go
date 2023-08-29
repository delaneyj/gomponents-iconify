package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Project(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v7h1V0H0zm7 0v7h1V0H7zM2 1v1h2V1H2zm1 2v1h2V3H3zm1 2v1h2V5H4z"/>`),
		g.Group(children),
	)
}