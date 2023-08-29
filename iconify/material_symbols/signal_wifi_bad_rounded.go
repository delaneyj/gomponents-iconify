package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiBadRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 19.925L17.1 21.3q-.275.275-.688.288T15.7 21.3q-.275-.275-.275-.7t.275-.7l1.4-1.4l-1.4-1.4q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.4 1.4l1.4-1.4q.275-.275.688-.287t.712.287q.275.275.275.7t-.275.7l-1.375 1.4l1.375 1.4q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275l-1.4-1.375Zm-7.2.375L.75 9.75q-.3-.3-.3-.738T.775 8.3Q3.1 6.175 5.988 5.087T12 4q3.05 0 6.025 1.088T23.25 8.3q.325.275.325.7t-.3.725L20.3 12.7q-.425-.125-.875-.2t-.925-.075q-2.525 0-4.3 1.763T12.425 18.5q0 .475.075.925t.2.875q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}