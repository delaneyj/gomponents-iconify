package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeOption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<rect width="18" height="18" x="3" y="3" fill="none" stroke="currentColor" stroke-width="2" rx="4"/>`),
		g.Group(children),
	)
}