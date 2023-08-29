package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-opacity=".5" stroke-width="2" d="M9 11h6v2H9v-2Z"/><path fill="currentColor" d="M19 7H5V5h14v2Zm0 12H5v-2h14v2Z"/></g>`),
		g.Group(children),
	)
}