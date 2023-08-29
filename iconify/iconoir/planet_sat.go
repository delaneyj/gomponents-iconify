package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanetSat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="8"/><path d="M17.5 6.348c2.297-.538 3.945-.476 4.338.312c.73 1.466-3.158 4.89-8.687 7.645c-5.528 2.757-10.602 3.802-11.333 2.336c-.392-.786.544-2.134 2.349-3.64"/><path stroke-linecap="round" stroke-linejoin="round" d="m9.5 10.51l.01-.011"/></g>`),
		g.Group(children),
	)
}