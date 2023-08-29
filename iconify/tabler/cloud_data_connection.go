package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDataConnection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9.897c0-1.714 1.46-3.104 3.26-3.104c.275-1.22 1.255-2.215 2.572-2.611c1.317-.397 2.77-.134 3.811.69c1.042.822 1.514 2.08 1.239 3.3h.693A2.42 2.42 0 0 1 19 10.586A2.42 2.42 0 0 1 16.575 13H8.26C6.46 13 5 11.61 5 9.897zM12 13v3m-2 2a2 2 0 1 0 4 0a2 2 0 1 0-4 0m4 0h7M3 18h7"/>`),
		g.Group(children),
	)
}