package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ufo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.94 10.28C15.66 7.87 14 6 12 6s-3.66 1.87-3.94 4.28C4.5 10.82 2 12.06 2 13.5C2 15.43 6.5 17 12 17s10-1.57 10-3.5c0-1.44-2.5-2.68-6.06-3.22Z"/>`),
		g.Group(children),
	)
}