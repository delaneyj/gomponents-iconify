package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Css(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 15q-.425 0-.713-.288T9.5 14v-1H11v.5h2v-1h-2.5q-.425 0-.713-.288T9.5 11.5V10q0-.425.288-.713T10.5 9h3q.425 0 .713.288T14.5 10v1H13v-.5h-2v1h2.5q.425 0 .713.288t.287.712V14q0 .425-.288.713T13.5 15h-3Zm6.5 0q-.425 0-.713-.288T16 14v-1h1.5v.5h2v-1H17q-.425 0-.713-.288T16 11.5V10q0-.425.288-.713T17 9h3q.425 0 .713.288T21 10v1h-1.5v-.5h-2v1H20q.425 0 .713.288T21 12.5V14q0 .425-.288.713T20 15h-3ZM4 15q-.425 0-.713-.288T3 14v-4q0-.425.288-.713T4 9h3q.425 0 .713.288T8 10v1H6.5v-.5h-2v3h2V13H8v1q0 .425-.288.713T7 15H4Z"/>`),
		g.Group(children),
	)
}