package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BikeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 18a3 3 0 1 0 6 0a3 3 0 1 0-6 0m14.437-1.56a3 3 0 0 0 4.123 4.123M22 18a3 3 0 0 0-3-3m-7 4v-4l-3-3l1.665-1.332m2.215-1.772L14 8l2 3h3m-3-6a1 1 0 1 0 2 0a1 1 0 1 0-2 0M3 3l18 18"/>`),
		g.Group(children),
	)
}