package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 4c2.2 0 4 1.8 4 4s-1.8 4-4 4s-4-1.8-4-4s1.8-4 4-4m7 17l1.8 1.77c.5.5 1.2.1 1.2-.49V18l2.8-3.4A1 1 0 0 0 22 13h-7c-.8 0-1.3 1-.8 1.6L17 18v3m-2-2.3l-2.3-2.8c-.4-.5-.6-1.1-.6-1.7c-.7-.2-1.4-.2-2.1-.2c-4.4 0-8 1.8-8 4v2h13v-1.3Z"/>`),
		g.Group(children),
	)
}