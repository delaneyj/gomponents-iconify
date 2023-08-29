package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoSensorThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 4h1a2 2 0 0 1 2 2v1m0 10v1a2 2 0 0 1-2 2h-1M7 20H6a2 2 0 0 1-2-2v-1M4 7V6a2 2 0 0 1 2-2h1m2 8a3 3 0 1 0 6 0a3 3 0 1 0-6 0m3 6v2m-8-8h2m6-8v2m8 6h-2"/>`),
		g.Group(children),
	)
}