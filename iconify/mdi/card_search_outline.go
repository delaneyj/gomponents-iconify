package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardSearchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 16c.87 0 1.69-.26 2.38-.7l2.44 2.44l1.42-1.42l-2.44-2.43A4.481 4.481 0 0 0 11.5 7C9 7 7 9 7 11.5S9 16 11.5 16m0-7a2.5 2.5 0 0 1 0 5a2.5 2.5 0 0 1 0-5M20 4H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2m0 14H4V6h16v12Z"/>`),
		g.Group(children),
	)
}