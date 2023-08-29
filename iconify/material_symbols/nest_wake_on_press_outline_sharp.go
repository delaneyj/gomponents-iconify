package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestWakeOnPressOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 13V4h3v9h-3ZM6.7 21L1 15.225l2.125-2.175l2.875.6V5.5q0-1.05.725-1.775T8.5 3q1.05 0 1.775.725T11 5.5V10h.85l5.375 2.375L15.75 21H6.7Zm.85-2h6.525L15 13.55l-4.25-1.875H9V5.5q0-.225-.138-.363T8.5 5q-.225 0-.363.138T8 5.5v10.6l-4.175-.875L7.55 19Zm0 0h6.525H7.55Z"/>`),
		g.Group(children),
	)
}