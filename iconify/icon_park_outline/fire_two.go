package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 21c0 7.18-5.82 13-13 13s-13-5.82-13-13s5-10 7-17c8 3.5 8 14 8 14s2-6 8-7.5c.5 5.5 3 7.434 3 10.5ZM7 36l32 8M7 44l32-8"/>`),
		g.Group(children),
	)
}