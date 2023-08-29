package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 6q-.625 0-1.063.438T6 7.5q0 .625.438 1.063T7.5 9q.625 0 1.063-.438T9 7.5q0-.625-.438-1.063T7.5 6Zm0 10q-.625 0-1.063.438T6 17.5q0 .625.438 1.063T7.5 19q.625 0 1.063-.438T9 17.5q0-.625-.438-1.063T7.5 16ZM4 3h16q.425 0 .713.288T21 4v7q0 .425-.288.713T20 12H4q-.425 0-.713-.288T3 11V4q0-.425.288-.713T4 3Zm1 2v5h14V5H5Zm-1 8h16q.425 0 .713.288T21 14v7q0 .425-.288.713T20 22H4q-.425 0-.713-.288T3 21v-7q0-.425.288-.713T4 13Zm1 2v5h14v-5H5ZM5 5v5v-5Zm0 10v5v-5Z"/>`),
		g.Group(children),
	)
}