package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormViewHide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 12h3c.5 2.5 3.273 5 6 5s5.5-2.5 6-5h3m-9 5v3m-4.5-4.5l-2 2m11-2l2 2"/>`),
		g.Group(children),
	)
}