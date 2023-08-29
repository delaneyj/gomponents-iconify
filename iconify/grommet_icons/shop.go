package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M4 7h16v16H4V7Zm4 2V5c0-2.21 1.795-4 4-4h0c2.21 0 4 1.795 4 4v4"/>`),
		g.Group(children),
	)
}