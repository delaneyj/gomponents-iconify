package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCarouselOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7h4v10H2V7Zm5-2h10v14H7V5Zm2 2v10V7Zm9 0h4v10h-4V7ZM9 7v10h6V7H9Z"/>`),
		g.Group(children),
	)
}