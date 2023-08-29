package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Javascript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 15q-.625 0-1.063-.438T6 13.5v-1h1.5v1H9V9h1.5v4.5q0 .625-.438 1.063T9 15H7.5Zm5.5 0q-.425 0-.713-.288T12 14v-1h1.5v.5h2v-1H13q-.425 0-.713-.288T12 11.5V10q0-.425.288-.713T13 9h3q.425 0 .713.288T17 10v1h-1.5v-.5h-2v1H16q.425 0 .713.288T17 12.5V14q0 .425-.288.713T16 15h-3Z"/>`),
		g.Group(children),
	)
}