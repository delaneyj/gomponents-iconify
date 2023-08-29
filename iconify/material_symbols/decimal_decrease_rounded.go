package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecimalDecreaseRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.525 19l.875.875q.275.3.288.713t-.288.712q-.3.3-.712.3t-.713-.3L12.4 18.7q-.3-.3-.3-.7t.3-.7l2.575-2.6q.3-.3.713-.3t.712.3q.3.3.288.713t-.288.712l-.875.875h5.15q.425 0 .713.288t.287.712q0 .425-.288.713t-.712.287h-5.15ZM4 13H3q-.425 0-.713-.288T2 12v-1q0-.425.288-.713T3 10h1q.425 0 .713.288T5 11v1q0 .425-.288.713T4 13Zm5.5 0q-1.45 0-2.475-1.025T6 9.5v-4q0-1.45 1.025-2.475T9.5 2q1.45 0 2.475 1.025T13 5.5v4q0 1.45-1.025 2.475T9.5 13Zm0-2q.625 0 1.063-.438T11 9.5v-4q0-.625-.438-1.063T9.5 4q-.625 0-1.063.438T8 5.5v4q0 .625.438 1.063T9.5 11Z"/>`),
		g.Group(children),
	)
}