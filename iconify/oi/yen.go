package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="m0 0l2.25 3H0v1h3v1H0v1h3v2h1V6h3V5H4V4h3V3H4.75L7 0H6L3.69 3h-.38L1 0H0z"/>`),
		g.Group(children),
	)
}