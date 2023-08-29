package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextToSpeech(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h9l-2 2H4v16h11v-3h2v3q0 .825-.588 1.413T15 22H4Zm2-4v-2h7v2H6Zm0-3v-2h5v2H6Zm9 0l-4-4H8V6h3l4-4v13Zm2-3.05v-6.9q.9.525 1.45 1.425T19 8.5q0 1.125-.55 2.025T17 11.95Zm0 4.3v-2.1q1.75-.625 2.875-2.163T21 8.5q0-1.95-1.125-3.488T17 2.85V.75q2.6.675 4.3 2.813T23 8.5q0 2.8-1.7 4.938T17 16.25Z"/>`),
		g.Group(children),
	)
}