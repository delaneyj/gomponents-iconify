package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDescending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M2 0v6H0l2.5 2L5 6H3V0H2zm2 0v1h4V0H4zm0 2v1h3V2H4zm0 2v1h2V4H4z"/>`),
		g.Group(children),
	)
}