package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 6q-.625 0-1.063.438T6 7.5q0 .625.438 1.063T7.5 9q.625 0 1.063-.438T9 7.5q0-.625-.438-1.063T7.5 6Zm0 10q-.625 0-1.063.438T6 17.5q0 .625.438 1.063T7.5 19q.625 0 1.063-.438T9 17.5q0-.625-.438-1.063T7.5 16ZM3 12V3h18v9H3Zm2-7v5h14V5H5ZM3 22v-9h18v9H3Zm2-7v5h14v-5H5ZM5 5v5v-5Zm0 10v5v-5Z"/>`),
		g.Group(children),
	)
}