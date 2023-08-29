package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricScooterSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 18q-1.25 0-2.125-.875T2 15q0-1.25.875-2.125T5 12q.975 0 1.738.563T7.8 14h5.3q.275-1.7 1.413-2.975T17.3 9.25L15.9 3H12V1h5.5l2.25 10H19q-1.65 0-2.825 1.175T15 15v1H7.8q-.3.875-1.062 1.438T5 18Zm14 0q-1.25 0-2.125-.875T16 15q0-1.25.875-2.125T19 12q1.25 0 2.125.875T22 15q0 1.25-.875 2.125T19 18Zm-6 5l-6-3h4v-2l6 3h-4v2Z"/>`),
		g.Group(children),
	)
}