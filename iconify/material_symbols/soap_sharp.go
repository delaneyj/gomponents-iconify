package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoapSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22H2V10.975l9.625-6.25L13.25 6.35L11.3 9.5H20v2h-8V13h10v2H12v1.5h9v2h-9V20h7v2ZM15.75 8q-.725 0-1.238-.513T14 6.25q0-.725.513-1.238T15.75 4.5q.725 0 1.238.513T17.5 6.25q0 .725-.513 1.238T15.75 8Zm1.75-4q-.625 0-1.063-.438T16 2.5q0-.625.438-1.063T17.5 1q.625 0 1.063.438T19 2.5q0 .625-.438 1.063T17.5 4ZM21 8q-.825 0-1.413-.588T19 6q0-.825.588-1.413T21 4q.825 0 1.413.588T23 6q0 .825-.588 1.413T21 8Z"/>`),
		g.Group(children),
	)
}