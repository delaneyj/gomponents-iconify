package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsChevronsRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M13.061 4.939l-2.122 2.122L15.879 12l-4.94 4.939l2.122 2.122L20.121 12z" fill="currentColor"/><path d="M6.061 19.061L13.121 12l-7.06-7.061l-2.122 2.122L8.879 12l-4.94 4.939z" fill="currentColor"/>`),
		g.Group(children),
	)
}