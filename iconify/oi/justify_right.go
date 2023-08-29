package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JustifyRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v1h8V0H0zm0 2v1h8V2H0zm0 2v1h8V4H0zm2 2v1h6V6H2z"/>`),
		g.Group(children),
	)
}