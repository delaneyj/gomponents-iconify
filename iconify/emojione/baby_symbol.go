package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabySymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#eda454"/><g fill="#fff"><path d="M37 19.5c0 1.2-1.2 2.5-2.5 2.5h-5c-1.2 0-2.5-1.2-2.5-2.5v-5c0-1.2 1.2-2.5 2.5-2.5h5c1.2 0 2.5 1.2 2.5 2.5v5M31 41l-5-3l-3 7l3 7l3-1l-2-6zm2 0l5-3l3 7l-3 7l-3-1l2-6z"/><path d="M36 24h-8.1L20 21l-1 2l7 5v6c0 2 1 3 3 3h6c2 0 3-1 3-3v-6l7-5l-1-2l-8 3"/></g>`),
		g.Group(children),
	)
}