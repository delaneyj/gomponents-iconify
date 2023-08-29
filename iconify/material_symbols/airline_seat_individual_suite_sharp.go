package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatIndividualSuiteSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 17V7h2v8h8V7h12v10H1Zm6-3q1.25 0 2.125-.875T10 11q0-1.25-.875-2.125T7 8q-1.25 0-2.125.875T4 11q0 1.25.875 2.125T7 14Z"/>`),
		g.Group(children),
	)
}