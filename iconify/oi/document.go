package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Document(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v8h7V4H3V0H0zm4 0v3h3L4 0zM1 2h1v1H1V2zm0 2h1v1H1V4zm0 2h4v1H1V6z"/>`),
		g.Group(children),
	)
}