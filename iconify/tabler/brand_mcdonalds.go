package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandMcdonalds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 20c0-3.952-.966-16-4.038-16S12 13.087 12 18.756C12 13.087 11.104 4 8.038 4C4.973 4 4 16.048 4 20"/>`),
		g.Group(children),
	)
}