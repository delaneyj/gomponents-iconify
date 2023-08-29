package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.42 7.72A7 7 0 0 0 5.06 9.61a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 13.5a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 12.5a5 5 0 0 0-3.58-4.78ZM12 12.83a2.94 2.94 0 0 0-3.43.53a2.93 2.93 0 0 0 0 4.13l2.72 2.72a1 1 0 0 0 1.42 0l2.72-2.72a2.93 2.93 0 0 0 0-4.13a2.94 2.94 0 0 0-3.43-.53Zm2 3.24l-2 2l-2-2a.88.88 0 0 1-.27-.65a.89.89 0 0 1 .27-.65a.92.92 0 0 1 1.3 0a1 1 0 0 0 1.42 0a.94.94 0 0 1 1.3 0a.89.89 0 0 1 .27.65a.88.88 0 0 1-.29.65Z"/>`),
		g.Group(children),
	)
}