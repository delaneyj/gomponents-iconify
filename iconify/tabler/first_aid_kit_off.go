package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAidKitOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.595 4.577A2 2 0 0 1 10 4h4a2 2 0 0 1 2 2v2m-4 0h6a2 2 0 0 1 2 2v6m-.576 3.405A2 2 0 0 1 18 20H6a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h2m2 6h4m-2-2v4M3 3l18 18"/>`),
		g.Group(children),
	)
}