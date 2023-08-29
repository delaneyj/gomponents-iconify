package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h18v18H1V1Zm4 18v4h18V5.97h-4M6 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM2 18l5-6l3 3l4-5l5 6"/>`),
		g.Group(children),
	)
}