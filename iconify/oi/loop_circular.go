package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoopCircular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 1C2.35 1 1 2.35 1 4H0l1.5 2L3 4H2c0-1.11.89-2 2-2V1zm2.5 1L5 4h1c0 1.11-.89 2-2 2v1c1.65 0 3-1.35 3-3h1L6.5 2z"/>`),
		g.Group(children),
	)
}