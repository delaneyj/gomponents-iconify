package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextToSpeechSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V2h11l-2 2H4v16h11v-3h2v5H2Zm4-4v-2h7v2H6Zm0-3v-2h5v2H6Zm9 0l-4-4H8V6h3l4-4v13Zm2-3.05v-6.9q.9.525 1.45 1.425T19 8.5q0 1.125-.55 2.025T17 11.95Zm0 4.3v-2.1q1.75-.625 2.875-2.163T21 8.5q0-1.95-1.125-3.488T17 2.85V.75q2.6.675 4.3 2.813T23 8.5q0 2.8-1.7 4.938T17 16.25Z"/>`),
		g.Group(children),
	)
}