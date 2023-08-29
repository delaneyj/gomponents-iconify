package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2.503 1.907A7.472 7.472 0 0 1 7.5 0a7.498 7.498 0 0 1 6.635 4H7.5a3.501 3.501 0 0 0-3.23 2.149L2.503 1.907Z"/><path fill="currentColor" d="M1.745 2.69a7.503 7.503 0 0 0 3.41 11.937l2.812-3.658a3.501 3.501 0 0 1-3.88-2.685a.502.502 0 0 1-.049-.092L1.745 2.69Z"/><path fill="currentColor" d="M6.215 14.89a7.5 7.5 0 0 0 8.357-9.895A.503.503 0 0 1 14.5 5H9.95A3.49 3.49 0 0 1 11 7.5a3.487 3.487 0 0 1-.954 2.405l-3.83 4.985Z"/><path fill="currentColor" d="M5 7.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0Z"/>`),
		g.Group(children),
	)
}