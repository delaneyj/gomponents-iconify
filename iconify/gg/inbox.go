package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v15a2 2 0 0 1-2 2H8.148a2.006 2.006 0 0 1-.148.005H4a2 2 0 0 1-2-2V5Zm3-1h14a1 1 0 0 1 1 1v9h-4a2 2 0 0 0-2 2v1h-4v-.995a2 2 0 0 0-2-2H4V5a1 1 0 0 1 1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}