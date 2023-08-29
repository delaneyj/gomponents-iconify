package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 4h-8a8 8 0 1 0 0 16h8a8 8 0 0 0-8-8a8 8 0 0 0 8-8M4 12h8m0-8v16"/>`),
		g.Group(children),
	)
}