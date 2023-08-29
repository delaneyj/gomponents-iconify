package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M44 24V9H4v30h20"/><circle cx="36" cy="34" r="5"/><path stroke-linecap="round" stroke-linejoin="round" d="m40 37l4 3M4 9l20 15L44 9"/></g>`),
		g.Group(children),
	)
}