package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RacetrackCycling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m8 1l3 3h1c2 0 2-3 0-3H8Zm.463 2a.5.5 0 0 0-.317.146l-2.5 2.5a.5.5 0 0 0 .042.745l2.158 1.726l-1.276 2.125a.501.501 0 1 0 .86.516l1.5-2.5a.5.5 0 0 0-.117-.649L7.304 6.402L9 4.707l1.146 1.147A.5.5 0 0 0 10.5 6H13a.5.5 0 0 0 0-1h-2.293L8.854 3.146A.5.5 0 0 0 8.463 3ZM2.949 7a3 3 0 1 0 .1 6a3 3 0 0 0-.099-6h-.001Zm9 0a3 3 0 1 0 .1 6a3 3 0 0 0-.099-6h-.001Zm.1 1a2 2 0 1 1-.098 3.998A2 2 0 0 1 12.05 8Z"/>`),
		g.Group(children),
	)
}