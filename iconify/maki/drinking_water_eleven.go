package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrinkingWaterEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5 11H3a.51.51 0 0 1-.5-.4L1 5.6a.5.5 0 0 1 .5-.6h5a.5.5 0 0 1 .5.6l-1.49 5A.51.51 0 0 1 5 11zM2.76 8h2.46l.67-2H2.11z" fill="currentColor"/><path d="M4.5 0A1.5 1.5 0 0 0 3 1.51v2a.5.5 0 0 0 .5.5h1A.5.5 0 0 0 5 3.5v-1a.5.5 0 0 1 .5-.5H10V0H4.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}