package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Update(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-1.875 0-3.513-.713t-2.85-1.924q-1.212-1.213-1.924-2.85T3 12q0-1.875.713-3.513t1.924-2.85q1.213-1.212 2.85-1.924T12 3q2.05 0 3.888.875T19 6.35V4h2v6h-6V8h2.75q-1.025-1.4-2.525-2.2T12 5Q9.075 5 7.037 7.038T5 12q0 2.925 2.038 4.963T12 19q2.625 0 4.588-1.7T18.9 13h2.05q-.375 3.425-2.937 5.713T12 21Zm2.8-4.8L11 12.4V7h2v4.6l3.2 3.2l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}