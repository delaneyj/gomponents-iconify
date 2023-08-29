package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.4 0L4 3.5H0l3 3L1.5 11l3.9-2.6L9.5 11L8 6.5l3-3H7L5.4 0z" fill="currentColor"/>`),
		g.Group(children),
	)
}