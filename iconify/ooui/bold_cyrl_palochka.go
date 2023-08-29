package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoldCyrlPalochka(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 19H6v-2l2-1V4L6 3V1h8v2l-2 1v12l2 1v2Z"/>`),
		g.Group(children),
	)
}