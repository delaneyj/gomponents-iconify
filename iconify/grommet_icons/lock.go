package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M19 23V11H5v12h14Zm-7-8v4m5-8V7c0-3 0-6-5-6S7 4 7 7v4"/>`),
		g.Group(children),
	)
}