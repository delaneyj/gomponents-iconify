package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalCarWashSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 5q-.625 0-1.063-.438T10.5 3.5q0-.525.363-1.125T12 1q.775.775 1.137 1.375T13.5 3.5q0 .625-.438 1.063T12 5ZM7 5q-.625 0-1.063-.438T5.5 3.5q0-.525.363-1.125T7 1q.775.775 1.137 1.375T8.5 3.5q0 .625-.438 1.063T7 5Zm10 0q-.625 0-1.063-.438T15.5 3.5q0-.525.363-1.125T17 1q.775.775 1.137 1.375T18.5 3.5q0 .625-.438 1.063T17 5ZM6 21v2H3v-9l2.45-7h13.1L21 14v9h-3v-2H6Zm-.2-9h12.4l-1.05-3H6.85L5.8 12Zm1.7 6q.625 0 1.063-.438T9 16.5q0-.625-.438-1.063T7.5 15q-.625 0-1.063.438T6 16.5q0 .625.438 1.063T7.5 18Zm9 0q.625 0 1.063-.438T18 16.5q0-.625-.438-1.063T16.5 15q-.625 0-1.063.438T15 16.5q0 .625.438 1.063T16.5 18Z"/>`),
		g.Group(children),
	)
}