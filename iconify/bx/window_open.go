package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 3H4c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h5v-2H4V7h16v12h-5v2h5c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2z"/><path fill="currentColor" d="M13 21v-5h3l-4-5l-4 5h3v5z"/>`),
		g.Group(children),
	)
}