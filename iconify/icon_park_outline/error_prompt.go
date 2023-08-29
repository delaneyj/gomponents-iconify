package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ErrorPrompt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m8 18l12 12m0-12L8 30M34 8c5 4.36 8 9.931 8 16c0 6.069-3 11.64-8 16m-7-26c3.75 2.726 6 6.207 6 10s-2.25 7.274-6 10"/>`),
		g.Group(children),
	)
}