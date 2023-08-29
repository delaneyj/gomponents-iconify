package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Export(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m16.95 5.968l-1.414 1.414L13 4.846v12.196h-2V4.847L8.465 7.382L7.05 5.968L12 1.018l4.95 4.95Z"/><path d="M5 20.982v-10h4v-2H3v14h18v-14h-6v2h4v10H5Z"/></g>`),
		g.Group(children),
	)
}