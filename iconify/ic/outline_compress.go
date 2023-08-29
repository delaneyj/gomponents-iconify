package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineCompress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 9v2h16V9H4zm12-5l-1.41-1.41L13 4.17V1h-2v3.19L9.39 2.61L8 4l4 4l4-4zM4 14h16v-2H4v2zm4 5l1.39 1.39L11 18.81V22h2v-3.17l1.59 1.59L16 19l-4-4l-4 4z"/>`),
		g.Group(children),
	)
}