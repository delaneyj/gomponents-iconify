package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricScooterOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 18q-1.25 0-2.125-.875T2 15q0-1.25.875-2.125T5 12q.975 0 1.738.563T7.8 14h5.3q.275-1.7 1.413-2.975T17.3 9.25L15.9 3H12V1h3.9q.675 0 1.238.438t.712 1.112l1.9 8.45H19q-1.65 0-2.825 1.175T15 15v1H7.8q-.3.875-1.062 1.438T5 18Zm0-2q.425 0 .713-.288T6 15q0-.425-.288-.713T5 14q-.425 0-.713.288T4 15q0 .425.288.713T5 16Zm14 2q-1.25 0-2.125-.875T16 15q0-1.25.875-2.125T19 12q1.25 0 2.125.875T22 15q0 1.25-.875 2.125T19 18Zm0-2q.425 0 .713-.288T20 15q0-.425-.288-.713T19 14q-.425 0-.713.288T18 15q0 .425.288.713T19 16Zm-6 7l-6-3h4v-2l6 3h-4v2Zm-8-8Zm14 0Z"/>`),
		g.Group(children),
	)
}