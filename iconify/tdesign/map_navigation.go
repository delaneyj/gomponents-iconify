package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapNavigation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2v3h10V2h2v3h2v2h-2v8h2v2h-2v4h-2v-4h-2v-2h2V7H8v4H6V7H2V5h4V2h2Zm6.58 8.418l-4.375 13.127l-3.008-5.743l-5.743-3.008l13.127-4.376Zm-8.032 4.786l2.13 1.117l1.117 2.13l1.624-4.87l-4.871 1.623Z"/>`),
		g.Group(children),
	)
}