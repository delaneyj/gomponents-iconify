package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Focus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 16a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm9 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path fill="currentColor" d="M9.5 16a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/></g>`),
		g.Group(children),
	)
}