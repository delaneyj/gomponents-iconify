package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaChangerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<g id="clarityMediaChangerLine0" fill="currentColor"><path d="M30 4H6a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h1.88v1.57a1 1 0 0 0 2 0V32h16v1.57a1 1 0 0 0 2 0V32H30a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2ZM6 30V6h24v24Z"/><path d="M20 18h2v2h-2zm4 0h2v2h-2zm-4 4h2v2h-2zm4 0h2v2h-2zm3.22-12H20v4a.8.8 0 1 0 1.59 0v-2.4h5.63a.8.8 0 1 0 0-1.6ZM8.81 10h8.14v2H8.81zm0 4h8.14v2H8.81zm0 4h8.14v2H8.81zm0 4h8.14v2H8.81zm0 4h8.14v2H8.81z"/></g>`),
		g.Group(children),
	)
}