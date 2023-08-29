package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HangerTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 9l-7.971 4.428A2 2 0 0 0 3 15.177V16a2 2 0 0 0 2 2h1"/><path d="M18 18h1a2 2 0 0 0 2-2v-.823a2 2 0 0 0-1.029-1.749L12 9c-1.457-.81-1.993-2.333-2-4a2 2 0 1 1 4 0"/><path d="M6 18a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2z"/></g>`),
		g.Group(children),
	)
}