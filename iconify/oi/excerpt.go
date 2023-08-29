package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Excerpt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v1h7V0H0zm0 2v1h5V2H0zm0 2v1h8V4H0zm0 2v1h1V6H0zm2 0v1h1V6H2zm2 0v1h1V6H4z"/>`),
		g.Group(children),
	)
}