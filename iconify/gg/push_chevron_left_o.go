package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushChevronLeftO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10 16H8V8h2v8Zm5.297-8.243l1.415 1.415L13.883 12l2.829 2.828l-1.415 1.415L11.055 12l4.242-4.243Z"/><path fill-rule="evenodd" d="M1 12c0 6.075 4.925 11 11 11s11-4.925 11-11S18.075 1 12 1S1 5.925 1 12Zm2 0a9 9 0 1 0 18 0a9 9 0 0 0-18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}