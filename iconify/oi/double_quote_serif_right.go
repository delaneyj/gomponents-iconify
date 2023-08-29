package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleQuoteSerifRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 1v3h2c0 1.11-.89 2-2 2v1c1.65 0 3-1.35 3-3V1H0zm5 0v3h2c0 1.11-.89 2-2 2v1c1.65 0 3-1.35 3-3V1H5z"/>`),
		g.Group(children),
	)
}