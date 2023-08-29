package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoldCyrlZhe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 10L1 3h4l3 7V3h4v7l3-7h4l-3.5 7l3.5 7h-4l-3-7v7H8v-7l-3 7H1l3.5-7Z"/>`),
		g.Group(children),
	)
}