package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRecognition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4c0-1.1.9-2 2-2h4v2H4v4H2V4m20 16c0 1.11-.89 2-2 2h-4v-2h4v-4h2v4M4 22a2 2 0 0 1-2-2v-4h2v4h4v2H4M20 2a2 2 0 0 1 2 2v4h-2V4h-4V2h4M9 7v2h2v8h2V9h2V7H9Z"/>`),
		g.Group(children),
	)
}