package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.51 1.51H2.49m16.02 18H2.49m4.006-5.992l4 4l4.015-4m-8.015-6l4-4l4.015 4M10.51 17.51v-14"/>`),
		g.Group(children),
	)
}