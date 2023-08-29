package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleQuoteSerifLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3 1C1.35 1 0 2.35 0 4v3h3V4H1c0-1.11.89-2 2-2V1zm5 0C6.35 1 5 2.35 5 4v3h3V4H6c0-1.11.89-2 2-2V1z"/>`),
		g.Group(children),
	)
}