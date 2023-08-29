package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm7-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill-rule="evenodd" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm0-2a8 8 0 0 0 7.634-10.4c-.835.226-1.713.346-2.619.346a9.996 9.996 0 0 1-8.692-5.053A8 8 0 0 0 12 20Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}