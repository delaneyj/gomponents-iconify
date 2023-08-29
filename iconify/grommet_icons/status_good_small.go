package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusGoodSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="6" cy="6" r="5" fill="currentColor" fill-rule="evenodd" stroke="currentColor"/>`),
		g.Group(children),
	)
}