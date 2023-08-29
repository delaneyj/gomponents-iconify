package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeceasedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22q-2.5 0-4.25-1.75T5 16v-2q1.775-.025 3.35.725T11 16.75v-3.825q-2.15-.35-3.575-2.013T6 7V1.425L9.5 4.45L12 1.425l2.5 3.025L18 1.425V7q0 2.25-1.425 3.913T13 12.925v3.825q1.075-1.275 2.65-2.025T19 14v2q0 2.5-1.75 4.25T13 22h-2Z"/>`),
		g.Group(children),
	)
}