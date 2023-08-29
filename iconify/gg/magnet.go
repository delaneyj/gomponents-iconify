package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-opacity=".5" d="M8 2.5H4v3h4v-3Zm12 0h-4v3h4v-3Z"/><path d="M8 7.5H4v6a8 8 0 1 0 16 0v-6h-4v6a4 4 0 0 1-8 0v-6Z"/></g>`),
		g.Group(children),
	)
}