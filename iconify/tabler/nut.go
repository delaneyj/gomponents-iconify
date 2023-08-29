package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 6.84a2.007 2.007 0 0 1 1 1.754v6.555c0 .728-.394 1.4-1.03 1.753l-6 3.844a1.995 1.995 0 0 1-1.94 0l-6-3.844A2.006 2.006 0 0 1 4 15.15V8.593c0-.728.394-1.399 1.03-1.753l6-3.582a2.049 2.049 0 0 1 2 0l6 3.582H19z"/><path d="M9 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/></g>`),
		g.Group(children),
	)
}