package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 5H2v4h2v2h2v2H4v2H2v4h14v-2h2v-2h2v-2h2v-2h-2V9h-2V7h-2V5zm0 2v2h2v2h2v2h-2v2h-2v2H4v-2h2v-2h2v-2H6V9H4V7h12z"/>`),
		g.Group(children),
	)
}