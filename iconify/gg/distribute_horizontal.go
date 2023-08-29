package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-opacity=".5" stroke-width="2" d="M11 9h2v6h-2V9Z"/><path fill="currentColor" d="M5 5v14h2V5H5Zm12 0v14h2V5h-2Z"/></g>`),
		g.Group(children),
	)
}