package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanetLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20 12a8 8 0 1 1-16 0a8 8 0 0 1 16 0Z"/><path d="M17.849 6.19c2.011-.37 3.49-.21 3.98.573c1.011 1.616-2.57 5.271-7.998 8.163c-5.429 2.893-10.649 3.927-11.66 2.31c-.533-.852.21-2.27 1.829-3.846"/></g>`),
		g.Group(children),
	)
}