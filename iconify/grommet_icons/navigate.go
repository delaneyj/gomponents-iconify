package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Navigate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m20 11l2-3l-2-3h-8v6h8Zm-8 13V0M4 2L2 5l2 3h8V2H4Z"/>`),
		g.Group(children),
	)
}