package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassCalibrationRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.538-1.463T7 17q0-2.075 1.463-3.538T12 12q2.075 0 3.538 1.463T17 17q0 2.075-1.463 3.538T12 22ZM6.275 11.425l-3.55-3.55q-.3-.3-.288-.713t.313-.687Q4.675 4.8 7.05 3.9T12 3q2.575 0 4.95.9t4.3 2.575q.3.275.313.688t-.288.712l-3.55 3.55q-.275.275-.7.3t-.775-.25q-.925-.725-2-1.1T12 10q-1.175 0-2.25.375t-2 1.1q-.35.275-.775.25t-.7-.3Z"/>`),
		g.Group(children),
	)
}