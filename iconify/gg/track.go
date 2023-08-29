package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Track(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path fill-rule="evenodd" d="M12 3a1 1 0 0 1 1 1v1.07A7.004 7.004 0 0 1 18.93 11H20a1 1 0 1 1 0 2h-1.07A7.004 7.004 0 0 1 13 18.93V20a1 1 0 1 1-2 0v-1.07A7.004 7.004 0 0 1 5.07 13H4a1 1 0 1 1 0-2h1.07A7.005 7.005 0 0 1 11 5.07V4a1 1 0 0 1 1-1Zm-5 9a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}