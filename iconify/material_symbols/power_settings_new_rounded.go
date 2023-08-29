package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSettingsNewRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13q-.425 0-.713-.288T11 12V4q0-.425.288-.713T12 3q.425 0 .713.288T13 4v8q0 .425-.288.713T12 13Zm0 8q-1.875 0-3.513-.713t-2.85-1.924q-1.212-1.213-1.924-2.85T3 12q0-1.725.625-3.3t1.8-2.85q.275-.3.7-.3t.725.3q.275.275.25.688t-.3.737q-.875.975-1.337 2.188T5 12q0 2.925 2.038 4.963T12 19q2.925 0 4.963-2.038T19 12q0-1.325-.475-2.563t-1.35-2.212q-.275-.3-.288-.7t.263-.675q.3-.3.725-.3t.7.3q1.175 1.275 1.8 2.85T21 12q0 1.875-.713 3.513t-1.924 2.85q-1.213 1.212-2.85 1.925T12 21Z"/>`),
		g.Group(children),
	)
}