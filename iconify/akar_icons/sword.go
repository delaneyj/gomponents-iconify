package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 3l4-1l-1 4l-10 10l-2.5-.5L8 13L18 3ZM2 20l2 2m1-8l1 4l4 1m-4-1l-3 3"/>`),
		g.Group(children),
	)
}