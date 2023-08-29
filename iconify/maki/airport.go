package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M15 6.818V8.5l-6.5-1l-.318 4.773L11 14v1l-3.5-.682L4 15v-1l2.818-1.727L6.5 7.5L0 8.5V6.818L6.5 4.5v-3s0-1.5 1-1.5s1 1.5 1 1.5v2.818l6.5 2.5z"/>`),
		g.Group(children),
	)
}