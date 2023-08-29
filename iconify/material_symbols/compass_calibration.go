package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassCalibration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.538-1.463T7 17q0-2.075 1.463-3.538T12 12q2.075 0 3.538 1.463T17 17q0 2.075-1.463 3.538T12 22Zm-5-9.85l-5-5q2-2 4.588-3.075T12 3q2.825 0 5.413 1.075T22 7.15l-5 5q-1.025-1.025-2.3-1.588T12 10q-1.425 0-2.7.563T7 12.15Z"/>`),
		g.Group(children),
	)
}