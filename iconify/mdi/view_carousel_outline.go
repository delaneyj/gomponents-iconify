package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCarouselOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6h4v11H2V6m5 13h10V4H7v15M9 6h6v11H9V6m9 0h4v11h-4V6Z"/>`),
		g.Group(children),
	)
}