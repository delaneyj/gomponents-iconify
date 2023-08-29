package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PediatricsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 7V5h4V2h2v3h4v2H7Zm2 15q-.825 0-1.413-.588T7 20v-2h5v-2H7v-2h5v-2H7v-1q0-1.25.875-2.125T10 8h4q1.25 0 2.125.875T17 11v9q0 .825-.588 1.413T15 22H9Z"/>`),
		g.Group(children),
	)
}