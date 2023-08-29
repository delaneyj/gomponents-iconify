package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlPointDuplicateOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 16h2v-3h3v-2h-3V8h-2v3h-3v2h3v3Zm-8 4.5q-2.725-.95-4.362-3.287T0 12q0-2.875 1.638-5.213T6 3.5v2.2q-1.85.875-2.925 2.575T2 12q0 2.025 1.075 3.725T6 18.3v2.2Zm9 .5q-1.875 0-3.513-.713t-2.85-1.924q-1.212-1.213-1.924-2.85T6 12q0-1.875.713-3.513t1.925-2.85q1.212-1.212 2.85-1.924T15 3q1.875 0 3.513.713t2.85 1.924q1.212 1.213 1.925 2.85T24 12q0 1.875-.713 3.513t-1.924 2.85q-1.213 1.212-2.85 1.925T15 21Zm0-9Zm0 7q2.925 0 4.963-2.038T22 12q0-2.925-2.038-4.963T15 5q-2.925 0-4.963 2.038T8 12q0 2.925 2.038 4.963T15 19Z"/>`),
		g.Group(children),
	)
}