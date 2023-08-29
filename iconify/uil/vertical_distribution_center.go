package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalDistributionCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 15h-1v-2a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v2H3a1 1 0 0 0 0 2h1v2a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-2h1a1 1 0 0 0 0-2Zm-3 3H6v-4h12ZM3 8h2v1a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V8h2a1 1 0 0 0 0-2h-2V5a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v1H3a1 1 0 0 0 0 2Zm4-2h10v2H7Z"/>`),
		g.Group(children),
	)
}