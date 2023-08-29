package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasTank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path stroke-width="1.493" d="M3 7.562A2.562 2.562 0 0 1 5.563 5H7V3h5v2h2.002A6.998 6.998 0 0 1 21 11.998v6.442a2.562 2.562 0 0 1-2.563 2.562H5.563A2.565 2.565 0 0 1 3 18.44V7.562Z" clip-rule="evenodd"/><path stroke-width="1.502" d="m8 8.878l8 8.238l-4-4.121l-4 4.12l4-4.12l4-4.121"/></g>`),
		g.Group(children),
	)
}