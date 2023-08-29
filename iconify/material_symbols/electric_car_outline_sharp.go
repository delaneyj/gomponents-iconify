package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricCarOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 15v2H3V8l2.45-7h13.1L21 8v9h-3v-2H6Zm-.2-9h12.4l-1.05-3H6.85L5.8 6ZM5 8v5v-5Zm2.5 4q.625 0 1.063-.438T9 10.5q0-.625-.438-1.063T7.5 9q-.625 0-1.063.438T6 10.5q0 .625.438 1.063T7.5 12Zm9 0q.625 0 1.063-.438T18 10.5q0-.625-.438-1.063T16.5 9q-.625 0-1.063.438T15 10.5q0 .625.438 1.063T16.5 12ZM13 23l-6-3h4v-2l6 3h-4v2ZM5 13h14V8H5v5Z"/>`),
		g.Group(children),
	)
}