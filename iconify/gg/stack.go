package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M20 4v12h2V2H8v2h12Z"/><path fill-rule="evenodd" d="M2 8v14h14V8H2Zm12 2H4v10h10V10Z" clip-rule="evenodd"/><path d="M17 7H5V5h14v14h-2V7Z"/></g>`),
		g.Group(children),
	)
}