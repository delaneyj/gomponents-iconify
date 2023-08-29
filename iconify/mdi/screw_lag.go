package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrewLag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 19.3l4-2.7V20l-2 2l-2-2v-.7m4-6.6l-4 2.7v2L9 18v1l6-3.9V14l-1 .7v-2M7 2v3h10V2H7m2 4v3l1 .7v3.7L9 14v1l6-3.9V10l-1 .7v-1l1-.7V6H9Z"/>`),
		g.Group(children),
	)
}