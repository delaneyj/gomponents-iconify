package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaJ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 7v8h-2v-1H9v1a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V7h-2Z"/>`),
		g.Group(children),
	)
}