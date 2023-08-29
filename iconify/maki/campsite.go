package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Campsite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 10v1a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1h.25l4.782-7.742a.55.55 0 0 1 .936 0L12.75 9H13a1 1 0 0 1 1 1Zm-3.5-1l-3-5l-3 5h6Z"/>`),
		g.Group(children),
	)
}