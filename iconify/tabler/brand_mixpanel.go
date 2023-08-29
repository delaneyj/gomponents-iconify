package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandMixpanel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12a2.5 2.5 0 1 0 5 0a2.5 2.5 0 1 0-5 0m17 0a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0m-8 0a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/>`),
		g.Group(children),
	)
}