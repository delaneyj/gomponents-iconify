package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlantOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17v2a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-4h8m-3.1-7.092a6 6 0 0 0-4.79-4.806M3 3v2a6 6 0 0 0 6 6h2m1.531-2.472A6 6 0 0 1 18 5h3v1a6 6 0 0 1-5.037 5.923M12 15v-3M3 3l18 18"/>`),
		g.Group(children),
	)
}