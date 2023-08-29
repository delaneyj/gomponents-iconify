package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 15a2 2 0 1 0-2-2l-1.9.87a2 2 0 0 0-1.1-.33a2 2 0 0 0 0 4a1.88 1.88 0 0 0 .92-.24l2.1 1a2 2 0 1 0 .8-1.84l-1.75-.8l1.91-.88a2 2 0 0 0 1.02.22Zm3.92-7.78A7 7 0 0 0 5.06 9.11a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 13a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 12a5 5 0 0 0-3.58-4.78Z"/>`),
		g.Group(children),
	)
}