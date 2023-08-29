package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nametag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 14v6h6v-2H6v-4H4Z"/><path fill-rule="evenodd" d="M9 9v6h6V9H9Zm4 2h-2v2h2v-2Z" clip-rule="evenodd"/><path d="M4 10V4h6v2H6v4H4Zm16 0V4h-6v2h4v4h2Zm0 4v6h-6v-2h4v-4h2Z"/></g>`),
		g.Group(children),
	)
}