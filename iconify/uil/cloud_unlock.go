package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUnlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 15.5h-3v-1a1 1 0 0 1 1.88-.5a1 1 0 0 0 1.37.34a1 1 0 0 0 .34-1.34a3.08 3.08 0 0 0-.46-.59A3 3 0 0 0 12 11.5a3 3 0 0 0-3 3v1.18a3 3 0 0 0 1 5.82h4a3 3 0 0 0 0-6Zm0 4h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2Zm4.42-12.78A7 7 0 0 0 5.06 8.61a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 12.5a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 11.5a5 5 0 0 0-3.58-4.78Z"/>`),
		g.Group(children),
	)
}