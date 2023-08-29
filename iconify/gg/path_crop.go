package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PathCrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="2" d="M6 6h8v8H6z" opacity=".5"/><path fill="currentColor" fill-rule="evenodd" d="M9 9h10v10H9V9Zm6 2h2v6h-6v-2h4v-4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}