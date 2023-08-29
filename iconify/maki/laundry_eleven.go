package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaundryEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5 0L4 2H2S1 2 1 3v7s0 1 1 1h7c1 0 1-1 1-1V1c0-1-1-1-1-1H5zm1 1h3v1H6V1zm-.5 3a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5z" fill="currentColor"/>`),
		g.Group(children),
	)
}