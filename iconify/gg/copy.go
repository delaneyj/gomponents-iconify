package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13 7H7V5h6v2Zm0 4H7V9h6v2Zm-6 4h6v-2H7v2Z"/><path fill-rule="evenodd" d="M3 19V1h14v4h4v18H7v-4H3Zm12-2V3H5v14h10Zm2-10v12H9v2h10V7h-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}