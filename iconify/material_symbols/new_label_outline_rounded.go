package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewLabelOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 19h-2q-.425 0-.713-.288T12 18q0-.425.288-.713T13 17h2l3.55-5L15 7H5v2q0 .425-.288.713T4 10q-.425 0-.713-.288T3 9V7q0-.825.588-1.413T5 5h10q.5 0 .938.225t.712.625l3.525 5q.375.525.375 1.15t-.375 1.15l-3.525 5q-.275.4-.712.625T15 19Zm-3.225-7ZM5 17H3q-.425 0-.713-.288T2 16q0-.425.288-.713T3 15h2v-2q0-.425.288-.713T6 12q.425 0 .713.288T7 13v2h2q.425 0 .713.288T10 16q0 .425-.288.713T9 17H7v2q0 .425-.288.713T6 20q-.425 0-.713-.288T5 19v-2Z"/>`),
		g.Group(children),
	)
}