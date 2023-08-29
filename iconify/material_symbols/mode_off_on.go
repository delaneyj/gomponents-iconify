package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeOffOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12q-.425 0-.713-.288T11 11V3q0-.425.288-.713T12 2q.425 0 .713.288T13 3v8q0 .425-.288.713T12 12Zm0 9q-1.875 0-3.513-.713t-2.85-1.924q-1.212-1.213-1.924-2.85T3 12q0-1.525.5-2.963T4.95 6.4q.275-.35.7-.338t.75.338q.275.275.25.675t-.275.75Q5.7 8.725 5.35 9.8T5 12q0 2.925 2.038 4.963T12 19q2.925 0 4.963-2.038T19 12q0-1.15-.338-2.238T17.6 7.775q-.25-.325-.275-.713t.25-.662q.3-.3.725-.313t.7.313q.975 1.2 1.488 2.625T21 12q0 1.875-.713 3.513t-1.924 2.85q-1.213 1.212-2.85 1.925T12 21Z"/>`),
		g.Group(children),
	)
}