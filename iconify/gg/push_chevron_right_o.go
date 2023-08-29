package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushChevronRightO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M14 8h2v8h-2V8Zm-5.297 8.243l-1.415-1.414L10.117 12L7.288 9.172l1.415-1.415L12.945 12l-4.242 4.243Z"/><path fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11Zm-2 0a9 9 0 1 0-18 0a9 9 0 0 0 18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}