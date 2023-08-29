package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanetBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M8 5.07A8 8 0 1 1 5.07 8"/><path stroke-linejoin="round" d="M13.83 14.926c-5.428 2.893-10.648 3.927-11.66 2.31c-.532-.852.211-2.27 1.83-3.846m13.849-7.2c2.011-.37 3.49-.21 3.98.573c.798 1.275-1.26 3.817-4.829 6.253"/></g>`),
		g.Group(children),
	)
}