package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RouteOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21V8.825Q4.125 8.5 3.562 7.738T3 6q0-1.25.875-2.125T6 3q1.25 0 2.125.875T9 6q0 .975-.563 1.738T7 8.825V19h4V3h8v12.175q.875.325 1.438 1.088T21 18q0 1.25-.875 2.125T18 21q-1.25 0-2.125-.875T15 18q0-.975.563-1.75T17 15.175V5h-4v16H5ZM6 7q.425 0 .713-.288T7 6q0-.425-.288-.713T6 5q-.425 0-.713.288T5 6q0 .425.288.713T6 7Zm12 12q.425 0 .713-.288T19 18q0-.425-.288-.713T18 17q-.425 0-.713.288T17 18q0 .425.288.713T18 19ZM6 6Zm12 12Z"/>`),
		g.Group(children),
	)
}