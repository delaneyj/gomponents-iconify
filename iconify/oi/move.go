package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3.5 0L2 1.5h1V3H1.5V2L0 3.5L1.5 5V4H3v1.5H2L3.5 7L5 5.5H4V4h1.5v1L7 3.5L5.5 2v1H4V1.5h1L3.5 0z"/>`),
		g.Group(children),
	)
}