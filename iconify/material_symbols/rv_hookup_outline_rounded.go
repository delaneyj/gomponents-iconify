package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RvHookupOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21q-.95 0-1.713-.55T6.2 19H5q-1.25 0-2.125-.875T2 16v-4q0-.825.588-1.413T4 10h4V7H3q-.425 0-.713-.288T2 6q0-.425.288-.713T3 5h12q.825 0 1.413.588T17 7v10h2.15l-.35-.35q-.3-.3-.288-.7t.288-.7q.3-.3.712-.313t.713.288L22.3 17.3q.3.3.3.7t-.3.7l-2.1 2.1q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l.4-.4h-7.4q-.3.9-1.075 1.45T9 21Zm1-11h5V7h-5v3Zm-1 9q.425 0 .713-.288T10 18q0-.425-.288-.713T9 17q-.425 0-.713.288T8 18q0 .425.288.713T9 19Zm-2.8-2q.325-.9 1.087-1.45T9 15q.95 0 1.725.55T11.8 17H15v-5H4v4q0 .425.288.713T5 17h1.2Zm0-5H4h11h-8.8Z"/>`),
		g.Group(children),
	)
}