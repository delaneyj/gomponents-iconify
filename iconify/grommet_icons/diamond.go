package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M6 3h12l4 6l-10 12L2 9l4-6ZM2 9h20M11 3L7 9l5 11m1-17l4 6l-5 11"/>`),
		g.Group(children),
	)
}