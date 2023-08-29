package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v2h7V0H0zm0 3v4.91c0 .05.04.09.09.09H6.9c.05 0 .09-.04.09-.09V3h-7zm1 1h1v1H1V4zm2 0h1v1H3V4zm2 0h1v1H5V4zM1 6h1v1H1V6zm2 0h1v1H3V6z"/>`),
		g.Group(children),
	)
}