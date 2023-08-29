package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AllInboxOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16h12v-3h-2.55q-.525.925-1.45 1.463T14 15q-1.05 0-1.975-.537T10.55 13H8v3Zm6-3q.85 0 1.425-.588T16 11h4V4H8v7h4q0 .825.588 1.413T14 13Zm-8 5V2h16v16H6Zm-4 4V6h2v14h14v2H2Zm6-6h12H8Z"/>`),
		g.Group(children),
	)
}