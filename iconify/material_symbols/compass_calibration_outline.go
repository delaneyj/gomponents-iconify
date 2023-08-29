package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassCalibrationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.538-1.463T7 17q0-2.075 1.463-3.538T12 12q2.075 0 3.538 1.463T17 17q0 2.075-1.463 3.538T12 22Zm0-2q1.25 0 2.125-.875T15 17q0-1.25-.875-2.125T12 14q-1.25 0-2.125.875T9 17q0 1.25.875 2.125T12 20Zm-5-7.85l-5-5q2-2 4.588-3.075T12 3q2.825 0 5.413 1.075T22 7.15l-5 5q-1.025-1.025-2.3-1.588T12 10q-1.425 0-2.7.563T7 12.15Zm.15-2.65q1.05-.75 2.288-1.125T12 8q1.325 0 2.538.363T16.85 9.45l2.2-2.2q-1.55-1.1-3.338-1.675T12 5q-1.925 0-3.713.575T4.95 7.25l2.2 2.25ZM12 8Zm0 9Z"/>`),
		g.Group(children),
	)
}