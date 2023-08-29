package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassCalibrationOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.538-1.463T7 17q0-2.075 1.463-3.538T12 12q2.075 0 3.538 1.463T17 17q0 2.075-1.463 3.538T12 22Zm0-2q1.25 0 2.125-.875T15 17q0-1.25-.875-2.125T12 14q-1.25 0-2.125.875T9 17q0 1.25.875 2.125T12 20Zm-5.725-8.575l-3.55-3.55q-.3-.3-.288-.713t.313-.687Q4.675 4.8 7.05 3.9T12 3q2.575 0 4.95.9t4.3 2.575q.3.275.313.688t-.288.712l-3.55 3.55q-.275.275-.7.3t-.775-.25q-.925-.725-2-1.1T12 10q-1.175 0-2.25.375t-2 1.1q-.35.275-.775.25t-.7-.3ZM7.15 9.5q1.05-.75 2.288-1.125T12 8q1.325 0 2.538.363T16.85 9.45l2.2-2.2q-1.55-1.1-3.338-1.675T12 5q-1.925 0-3.713.575T4.95 7.25l2.2 2.25ZM12 8Zm0 9Z"/>`),
		g.Group(children),
	)
}