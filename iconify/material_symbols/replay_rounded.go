package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplayRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-1.875 0-3.513-.713t-2.85-1.924q-1.212-1.213-1.924-2.85T3 13q0-.425.288-.713T4 12q.425 0 .713.288T5 13q0 2.925 2.038 4.963T12 20q2.925 0 4.963-2.038T19 13q0-2.925-2.038-4.963T12 6h-.15l.85.85q.3.3.288.7t-.288.7q-.3.3-.712.313t-.713-.288L8.7 5.7q-.3-.3-.3-.7t.3-.7l2.575-2.575q.3-.3.713-.288t.712.313q.275.3.288.7t-.288.7l-.85.85H12q1.875 0 3.513.713t2.85 1.925q1.212 1.212 1.925 2.85T21 13q0 1.875-.713 3.513t-1.924 2.85q-1.213 1.212-2.85 1.925T12 22Z"/>`),
		g.Group(children),
	)
}