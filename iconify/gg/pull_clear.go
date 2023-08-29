package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PullClear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 6H2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6h-2v10H4V6Z"/><path d="M6 12h12v2H6v-2Zm12-4H6v2h12V8Z"/></g>`),
		g.Group(children),
	)
}