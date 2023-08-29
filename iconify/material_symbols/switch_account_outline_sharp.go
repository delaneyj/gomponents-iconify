package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchAccountOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 11q1.25 0 2.125-.875T17 8q0-1.25-.875-2.125T14 5q-1.25 0-2.125.875T11 8q0 1.25.875 2.125T14 11Zm-6 4.75q1.125-1.325 2.7-2.038T14 13q1.725 0 3.3.713T20 15.75V4H8v11.75ZM6 18V2h16v16H6Zm-4 4V6h2v14h14v2H2ZM14 9q-.425 0-.713-.288T13 8q0-.425.288-.713T14 7q.425 0 .713.288T15 8q0 .425-.288.713T14 9Zm-3.3 7h6.6q-.725-.5-1.575-.75T14 15q-.875 0-1.725.25T10.7 16ZM14 9.875Z"/>`),
		g.Group(children),
	)
}