package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageSquareX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 2H8C4.691 2 2 4.691 2 8v13a1 1 0 0 0 1 1h13c3.309 0 6-2.691 6-6V8c0-3.309-2.691-6-6-6zm.706 13.293l-1.414 1.414L12 13.415l-3.292 3.292l-1.414-1.414l3.292-3.292l-3.292-3.292l1.414-1.414L12 10.587l3.292-3.292l1.414 1.414l-3.292 3.292l3.292 3.292z"/>`),
		g.Group(children),
	)
}