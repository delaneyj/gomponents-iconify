package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chinese(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" stroke-linejoin="round" rx="3"/><path stroke-linejoin="round" d="M14 18h20v10H14z"/><path d="M24 14v21"/></g>`),
		g.Group(children),
	)
}