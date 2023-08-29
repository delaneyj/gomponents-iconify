package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M.09 0C.03 0 0 .04 0 .09V7.9c0 .05.04.09.09.09H6.9c.05 0 .09-.04.09-.09V.09C6.99.03 6.95 0 6.9 0H.09zM1 1h5v2H1V1zm0 3h1v1H1V4zm2 0h1v1H3V4zm2 0h1v3H5V4zM1 6h1v1H1V6zm2 0h1v1H3V6z"/>`),
		g.Group(children),
	)
}