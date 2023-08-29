package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 3H4c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2zM4 19V7h16l.001 12H4z"/><path fill="currentColor" d="m15.707 10.707l-1.414-1.414L12 11.586L9.707 9.293l-1.414 1.414L10.586 13l-2.293 2.293l1.414 1.414L12 14.414l2.293 2.293l1.414-1.414L13.414 13z"/>`),
		g.Group(children),
	)
}