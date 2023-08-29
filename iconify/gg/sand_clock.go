package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SandClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13 6h-2v1a1 1 0 1 0 2 0V6Z"/><path fill-rule="evenodd" d="M6 2v2h1v3a5 5 0 0 0 5 5a5 5 0 0 0-5 5v3H6v2h12v-2h-1v-3a5 5 0 0 0-5-5a5 5 0 0 0 5-5V4h1V2H6Zm3 2h6v3a3 3 0 1 1-6 0V4Zm0 13v3h6v-3a3 3 0 1 0-6 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}