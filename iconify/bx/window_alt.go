package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 3H4c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2zm0 2l.001 4H4V5h16zM4 19v-8h16.001l.001 8H4z"/><path fill="currentColor" d="M14 6h2v2h-2zm3 0h2v2h-2z"/>`),
		g.Group(children),
	)
}