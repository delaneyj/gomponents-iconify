package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Display(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="36" height="20" x="6" y="6" rx="2"/><path stroke-linecap="round" d="M14 13h8m-8 6h20M8 44l4.889-6h21.778L40 44M24 26v18"/></g>`),
		g.Group(children),
	)
}