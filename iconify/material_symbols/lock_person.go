package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockPerson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 8h6V6q0-1.25-.875-2.125T12 3q-1.25 0-2.125.875T9 6v2Zm9 15q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Zm-5.75-1H6q-.825 0-1.413-.588T4 20V10q0-.825.588-1.413T6 8h1V6q0-2.075 1.463-3.538T12 1q2.075 0 3.538 1.463T17 6v2h1q.825 0 1.413.588T20 10v1.3q-.5-.175-1-.238T18 11q-2.925 0-4.963 2.038T11 18q0 1.075.338 2.087T12.25 22ZM18 18q.625 0 1.063-.438T19.5 16.5q0-.625-.438-1.063T18 15q-.625 0-1.063.438T16.5 16.5q0 .625.438 1.063T18 18Zm0 3q.75 0 1.4-.35t1.075-.975q-.575-.35-1.2-.513T18 19q-.65 0-1.275.162t-1.2.513q.425.625 1.075.975T18 21Z"/>`),
		g.Group(children),
	)
}