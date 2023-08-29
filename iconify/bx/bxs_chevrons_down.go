package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsChevronsDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M16.939 10.939L12 15.879l-4.939-4.94l-2.122 2.122L12 20.121l7.061-7.06z" fill="currentColor"/><path d="M16.939 3.939L12 8.879l-4.939-4.94l-2.122 2.122L12 13.121l7.061-7.06z" fill="currentColor"/>`),
		g.Group(children),
	)
}