package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" stroke="currentColor"><path fill="currentColor" d="M11.5 18a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm10 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 21h-2m12.708-4.717a9 9 0 1 0-17.66-3.218A7.001 7.001 0 0 0 10 27h13a6 6 0 0 0 3.708-10.717Z"/></g>`),
		g.Group(children),
	)
}