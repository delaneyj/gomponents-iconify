package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestWakeOnPressOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 13q-.625 0-1.063-.438T19 11.5v-6q0-.625.438-1.063T20.5 4q.625 0 1.063.438T22 5.5v6q0 .625-.438 1.063T20.5 13ZM7.525 21q-.425 0-.775-.163t-.625-.437L1 15.225L2.4 13.8q.35-.35.8-.5t.925-.05l1.875.4V5.5q0-1.05.725-1.775T8.5 3q1.05 0 1.775.725T11 5.5V10h.65q.125 0 .225.025t.225.075l3.8 1.675q.6.275.9.875t.175 1.25l-.925 5.425q-.125.725-.688 1.2T14.076 21h-6.55Zm.025-2h6.525L15 13.55l-4.25-1.875H9V5.5q0-.225-.138-.363T8.5 5q-.225 0-.363.138T8 5.5v10.6l-4.175-.875L7.55 19Zm0 0h6.525H7.55Z"/>`),
		g.Group(children),
	)
}