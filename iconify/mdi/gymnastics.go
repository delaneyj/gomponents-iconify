package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gymnastics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6c0-1.1.9-2 2-2s2 .9 2 2s-.9 2-2 2s-2-.9-2-2M1 9h6l7-5l1.31 1.5l-4.17 3H14L21.8 4L23 5.4L14.5 12L14 22h-2l-.5-10L8 11H1V9Z"/>`),
		g.Group(children),
	)
}