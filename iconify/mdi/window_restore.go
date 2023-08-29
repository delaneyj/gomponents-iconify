package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowRestore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 8h4V4h12v12h-4v4H4V8m12 0v6h2V6h-8v2h6M6 12v6h8v-6H6Z"/>`),
		g.Group(children),
	)
}