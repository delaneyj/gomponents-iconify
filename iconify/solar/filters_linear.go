package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiltersLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M18 8A6 6 0 1 1 6 8a6 6 0 0 1 12 0Z"/><path d="M6.5 10.189a6 6 0 1 0 7.106 3.669"/><path d="M12 20.472a6 6 0 1 0 5.5-10.283"/></g>`),
		g.Group(children),
	)
}