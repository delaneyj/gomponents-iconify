package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StressManagement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16q1.7 0 3.775-.675T22 12.95v-.375q0-1.275-1.025-2.05t-2.25-.425q-.425.125-.863.3t-1.162.45q-1.525.65-2.575.9T12 12q-1.125 0-2.175-.262t-2.65-.913q-.6-.25-1.037-.412t-.838-.288q-1.225-.4-2.262.375T2 12.575v.325q3.275 1.5 5.75 2.3T12 16Zm0 6q3.45 0 6.063-1.888t3.537-4.837q-3.275 1.4-5.5 2.063T12 18q-1.9 0-4.338-.713T2.35 15.25q.85 3.075 3.425 4.913T12 22Zm0-12q-1.65 0-2.825-1.175T8 6q0-1.65 1.175-2.825T12 2q1.65 0 2.825 1.175T16 6q0 1.65-1.175 2.825T12 10Z"/>`),
		g.Group(children),
	)
}