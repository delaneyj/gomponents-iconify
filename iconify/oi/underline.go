package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M1 0v4c0 1.1 1.12 2 2.5 2H4c1.1 0 2-.9 2-2V0H5v4c0 .55-.45 1-1 1s-1-.45-1-1V0H1zM0 7v1h7V7H0z"/>`),
		g.Group(children),
	)
}