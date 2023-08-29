package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func External(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M15.64 7.025h-3.622v-2h7v7h-2v-3.55l-4.914 4.914l-1.414-1.414l4.95-4.95Z"/><path d="M10.982 6.975h-6v12h12v-6h-2v4h-8v-8h4v-2Z"/></g>`),
		g.Group(children),
	)
}