package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesWearablesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 18q1.45 0 2.475-1.025T21 14.5q0-1.45-1.025-2.475T17.5 11q-1.45 0-2.475 1.025T14 14.5q0 1.45 1.025 2.475T17.5 18ZM2 22V2h12v5.85q-.575.3-1.075.688T12 9.4V7H4v10h6.425q.425 1.275 1.275 2.3t2.025 1.7v1H2Zm13 0v-2.6q-1.4-.7-2.2-2.012T12 14.5q0-1.575.8-2.888T15 9.6V7h5v2.6q1.4.7 2.2 2.013T23 14.5q0 1.575-.8 2.888T20 19.4V22h-5Z"/>`),
		g.Group(children),
	)
}