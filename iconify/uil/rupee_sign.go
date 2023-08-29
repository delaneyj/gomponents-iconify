package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RupeeSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 7h-2.21a5.49 5.49 0 0 0-1-2H18a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2h3.5a3.5 3.5 0 0 1 3.15 2H7a1 1 0 0 0 0 2h7a3.5 3.5 0 0 1-3.45 3H7a.7.7 0 0 0-.14 0a.65.65 0 0 0-.2 0a.69.69 0 0 0-.19.1l-.12.07a.75.75 0 0 0-.14.17a1.1 1.1 0 0 0-.09.14a.61.61 0 0 0 0 .18A.65.65 0 0 0 6 13a.7.7 0 0 0 0 .14a.65.65 0 0 0 0 .2a.69.69 0 0 0 .1.19s0 .08.07.12l6 7a1 1 0 0 0 1.52-1.3L9.18 14h1.32A5.5 5.5 0 0 0 16 9h2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}