package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleQuoteSansLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 1v6l3-3V1H0zm5 0v6l3-3V1H5z"/>`),
		g.Group(children),
	)
}