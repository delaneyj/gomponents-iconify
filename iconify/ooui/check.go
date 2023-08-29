package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 14.2L2.8 10l-1.4 1.4L7 17L19 5l-1.4-1.4z"/>`),
		g.Group(children),
	)
}