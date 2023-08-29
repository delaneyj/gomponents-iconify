package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M16 6h16v6H16V6Z"/><path d="M6 42V6"/><path stroke-linejoin="round" d="M16 21h20v6H16v-6Zm0 15h26v6H16v-6Z"/></g>`),
		g.Group(children),
	)
}