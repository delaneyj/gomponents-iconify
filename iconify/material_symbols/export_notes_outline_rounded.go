package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExportNotesOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 17.7v1.8q0 .2.15.35t.35.15q.2 0 .35-.15t.15-.35v-3q0-.2-.15-.35T19.5 16h-3q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15h1.8l-2.45 2.45q-.15.15-.15.35t.15.35q.15.15.35.15t.35-.15L19 17.7ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v6.7q-.475-.225-.975-.388T19 11.075V5H5v14h6.05q.075.55.238 1.05t.387.95H5Zm0-3v1V5v6.075V11v7Zm2-2q0 .425.288.713T8 17h3.075q.075-.525.238-1.025t.362-.975H8q-.425 0-.713.288T7 16Zm0-4q0 .425.288.713T8 13h5.1q.8-.75 1.788-1.25T17 11.075q-.225-.05-.5-.063T16 11H8q-.425 0-.713.288T7 12Zm0-4q0 .425.288.713T8 9h8q.425 0 .713-.288T17 8q0-.425-.288-.713T16 7H8q-.425 0-.713.288T7 8Zm11 15q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Z"/>`),
		g.Group(children),
	)
}