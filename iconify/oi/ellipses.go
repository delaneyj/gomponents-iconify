package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ellipses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 3v2h2V3H0zm3 0v2h2V3H3zm3 0v2h2V3H6z"/>`),
		g.Group(children),
	)
}