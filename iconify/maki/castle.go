package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Castle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M11 4H4a1 1 0 0 1-1-1V.5a.5.5 0 0 1 1 0V2h1V1a1 1 0 1 1 2 0v1h1V1a1 1 0 1 1 2 0v1h1V.5a.5.5 0 0 1 1 0V3a1 1 0 0 1-1 1zm3 10.5a.5.5 0 0 1-.5.5h-12a.5.5 0 0 1 0-1H2a1 1 0 0 0 1-1s1-6 1-7a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1c0 1 1 7 1 7a1 1 0 0 0 1 1h.5a.5.5 0 0 1 .5.49v.01zm-5-4a1.5 1.5 0 0 0-3 0V14h3v-3.5z"/>`),
		g.Group(children),
	)
}