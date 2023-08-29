package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MisalignedSemicircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M20 30c-6.627 0-12-5.373-12-12S13.373 6 20 6"/><path d="M20 22a4 4 0 0 1 0-8m8 20a4 4 0 0 0 0-8"/><path d="M28 42c6.627 0 12-5.373 12-12s-5.373-12-12-12"/></g>`),
		g.Group(children),
	)
}