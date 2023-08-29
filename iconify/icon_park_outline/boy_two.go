package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="10" r="6"/><path d="M31 44v-9l5-3l-4-13s-4-3-8-3s-8 3-8 3l-4 12l5 4v9"/></g>`),
		g.Group(children),
	)
}