package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandPlanetscale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.993 11.63a9 9 0 0 1-9.362 9.362l9.362-9.362zM12 3a9.001 9.001 0 0 1 8.166 5.211L8.211 20.166A9 9 0 0 1 12 3zm0 9l-6 6"/>`),
		g.Group(children),
	)
}