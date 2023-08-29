package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M9 4a3 3 0 1 1 6 0v8a3 3 0 1 1-6 0V4Zm4 0v8a1 1 0 1 1-2 0V4a1 1 0 1 1 2 0Z" clip-rule="evenodd"/><path d="M18 12a6.002 6.002 0 0 1-5 5.917V21h4v2H7v-2h4v-3.083A6.002 6.002 0 0 1 6 12V9h2v3a4 4 0 0 0 8 0V9h2v3Z"/></g>`),
		g.Group(children),
	)
}