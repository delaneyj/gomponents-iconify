package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandStackshare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 6a2 2 0 1 0 4 0a2 2 0 1 0-4 0m0 12a2 2 0 1 0 4 0a2 2 0 1 0-4 0M3 12a2 2 0 1 0 4 0a2 2 0 1 0-4 0m4 0h3l3.5 6H17m0-12h-3.5L10 12"/>`),
		g.Group(children),
	)
}