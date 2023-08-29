package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeechToTextSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 10q-.825 0-1.413-.575T15 8V4q0-.85.588-1.425T17 2q.85 0 1.425.575T19 4v4q0 .85-.575 1.425T17 10ZM3 22V2h10v2H5v16h11v-2h2v4H3Zm4-4v-2h7v2H7Zm0-3v-2h5v2H7Zm11 1h-2v-2.6q-1.925-.35-3.213-1.863T11.5 8h2q0 1.45 1.025 2.475T17 11.5q1.475 0 2.488-1.025T20.5 8h2q0 2.025-1.275 3.538T18 13.4V16Z"/>`),
		g.Group(children),
	)
}