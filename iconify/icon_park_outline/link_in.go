package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 12V4h16v8m0 24v8H16v-8m2-12H4m40 0H30m-6 10V14m-11 5l5 5l-5 5m22-10l-5 5l5 5"/>`),
		g.Group(children),
	)
}