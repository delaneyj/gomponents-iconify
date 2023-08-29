package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlurMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17q-.425 0-.713-.288T7 16q0-.425.288-.713T8 15h3q-.35-.425-.563-.938T10.1 13H5.5q-.425 0-.713-.288T4.5 12q0-.425.288-.713T5.5 11h4.6q.125-.55.338-1.063T11 9H5q-.425 0-.713-.288T4 8q0-.425.288-.713T5 7h10q2.075 0 3.538 1.463T20 12q0 2.075-1.463 3.538T15 17H8Zm7-2q1.25 0 2.125-.875T18 12q0-1.25-.875-2.125T15 9q-1.25 0-2.125.875T12 12q0 1.25.875 2.125T15 15ZM5 17q-.425 0-.713-.288T4 16q0-.425.288-.713T5 15q.425 0 .713.288T6 16q0 .425-.288.713T5 17Z"/>`),
		g.Group(children),
	)
}