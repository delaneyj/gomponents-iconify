package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingBridge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 5v14M18 5v14M2 15h20M3 8a7.5 7.5 0 0 0 3-2a6.5 6.5 0 0 0 12 0a7.5 7.5 0 0 0 3 2m-9 2v5"/>`),
		g.Group(children),
	)
}