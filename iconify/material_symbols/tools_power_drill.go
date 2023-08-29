package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsPowerDrill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.75 10h4.5q.3 0 .525-.225T12 9.25q0-.3-.225-.525T11.25 8.5h-4.5q-.3 0-.525.225T6 9.25q0 .3.225.525T6.75 10Zm0-2.5h4.5q.3 0 .525-.225T12 6.75q0-.3-.225-.525T11.25 6h-4.5q-.3 0-.525.225T6 6.75q0 .3.225.525t.525.225ZM17 11V5h1q.825 0 1.413.588T20 7h2q.425 0 .713.288T23 8q0 .425-.288.713T22 9h-2q0 .825-.588 1.413T18 11h-1Zm-5 4.5H6V13q-1.65 0-2.825-1.175T2 9V7q0-1.65 1.175-2.825T6 3h8q.825 0 1.413.588T16 5v6q0 .825-.588 1.413T14 13h-2v2.5ZM5.5 21q-.625 0-1.063-.438T4 19.5V18q0-.625.438-1.063T5.5 16.5h7q.625 0 1.063.438T14 18v1.5q0 .625-.438 1.063T12.5 21h-7Z"/>`),
		g.Group(children),
	)
}