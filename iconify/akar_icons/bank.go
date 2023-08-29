package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 11V8l-8-5l-8 5v3h16ZM3 21h18M5 20v-5h2m10 5v-5h2m-8 5v-5h2"/>`),
		g.Group(children),
	)
}