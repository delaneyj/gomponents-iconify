package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v18l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2m-5.4 12L12 12.4L9.4 14l.7-3l-2.3-2l3-.3L12 6l1.2 2.8l3 .3l-2.3 2l.7 2.9Z"/>`),
		g.Group(children),
	)
}