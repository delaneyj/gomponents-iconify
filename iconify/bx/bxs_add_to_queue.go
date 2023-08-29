package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsAddToQueue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M4 22h12v-2H4V8H2v12c0 1.103.897 2 2 2z" fill="currentColor"/><path d="M20 2H8c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2zm-2 9h-3v3h-2v-3h-3V9h3V6h2v3h3v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}