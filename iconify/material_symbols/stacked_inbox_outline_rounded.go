package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedInboxOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17q-.825 0-1.413-.588T5 15V5q0-.825.588-1.413T7 3h14q.825 0 1.413.588T23 5v10q0 .825-.588 1.413T21 17H7Zm0-5v3h14v-3h-3.55q-.525.9-1.425 1.45T14 14q-1.125 0-2.025-.55T10.55 12H7Zm7 0q.65 0 1.175-.388t.725-.987q.05-.275.25-.45t.475-.175H21V5H7v5h4.375q.275 0 .475.175t.25.45q.2.6.725.988T14 12ZM3 21q-.825 0-1.413-.588T1 19V8q0-.425.288-.713T2 7q.425 0 .713.288T3 8v11h15q.425 0 .713.288T19 20q0 .425-.288.713T18 21H3Zm4-6h14H7Z"/>`),
		g.Group(children),
	)
}