package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forwardburger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 13H3v-2h16l-4-4l1.4-1.4l6.4 6.4l-6.4 6.4L15 17l4-4M3 6h10v2H3V6m10 10v2H3v-2h10Z"/>`),
		g.Group(children),
	)
}