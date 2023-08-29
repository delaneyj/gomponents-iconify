package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanetRingTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="7"/><path d="M18 7.96c2.59-.125 4.379.274 4.625 1.193c.429 1.6-3.98 4.172-9.849 5.745c-5.868 1.572-10.972 1.55-11.401-.051c-.254-.948 1.188-2.236 3.625-3.455"/></g>`),
		g.Group(children),
	)
}