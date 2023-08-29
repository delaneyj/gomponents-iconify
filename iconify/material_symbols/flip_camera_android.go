package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipCameraAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-3.575 0-6.325-2.25T2.2 14h2.05q.7 2.65 2.85 4.325T12 20q2.15 0 4-1.063T18.9 16H16v-2h6v6h-2v-2q-1.425 1.9-3.525 2.95T12 22Zm0-7q-1.25 0-2.125-.875T9 12q0-1.25.875-2.125T12 9q1.25 0 2.125.875T15 12q0 1.25-.875 2.125T12 15ZM2 10V4h2v2q1.425-1.9 3.525-2.95T12 2q3.575 0 6.325 2.25T21.8 10h-2.05q-.7-2.65-2.85-4.325T12 4Q9.85 4 8 5.063T5.1 8H8v2H2Z"/>`),
		g.Group(children),
	)
}