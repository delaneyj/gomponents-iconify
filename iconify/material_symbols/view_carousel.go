package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCarousel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7h4v10H2V7Zm5-2h10v14H7V5Zm11 2h4v10h-4V7Z"/>`),
		g.Group(children),
	)
}