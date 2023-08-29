package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jump(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h7.09v2H5v14h14v-5.09h2V21H3Zm7.586-9l7-7H13V3h8v8h-2V6.414l-7 7L10.586 12Z"/>`),
		g.Group(children),
	)
}