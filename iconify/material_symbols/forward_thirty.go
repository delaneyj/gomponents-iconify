package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 16v-1.5H10v-1H8.5v-1H10v-1H7.5V10h3q.425 0 .713.288T11.5 11v4q0 .425-.288.713T10.5 16h-3Zm6 0q-.425 0-.713-.288T12.5 15v-4q0-.425.288-.713T13.5 10h2q.425 0 .713.288T16.5 11v4q0 .425-.288.713T15.5 16h-2Zm.5-1.5h1v-3h-1v3ZM12 22q-1.875 0-3.513-.713t-2.85-1.924q-1.212-1.213-1.924-2.85T3 13q0-1.875.713-3.513t1.924-2.85q1.213-1.212 2.85-1.924T12 4h.15L10.6 2.45L12 1l4 4l-4 4l-1.4-1.45L12.15 6H12Q9.075 6 7.037 8.038T5 13q0 2.925 2.038 4.963T12 20q2.925 0 4.963-2.038T19 13h2q0 1.875-.713 3.513t-1.924 2.85q-1.213 1.212-2.85 1.925T12 22Z"/>`),
		g.Group(children),
	)
}