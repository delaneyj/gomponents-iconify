package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpDrafts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.99 6.86L12 1L2 6.86V20h20l-.01-13.14zM12 13L3.74 7.84L12 3l8.26 4.84L12 13z"/>`),
		g.Group(children),
	)
}