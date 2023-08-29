package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-opacity=".5" d="M13 10h4v6h-4v-6Z"/><path d="M11 4H7v12h4V4Zm7 14H6v2h12v-2Z"/></g>`),
		g.Group(children),
	)
}