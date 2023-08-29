package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LooksThreeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 17q.825 0 1.413-.588T15 15v-1.5q0-.65-.425-1.075T13.5 12q.65 0 1.075-.425T15 10.5V9q0-.825-.588-1.413T13 7h-3q-.425 0-.713.288T9 8q0 .425.288.713T10 9h3v2h-1q-.425 0-.713.288T11 12q0 .425.288.713T12 13h1v2h-3q-.425 0-.713.288T9 16q0 .425.288.713T10 17h3Zm-8 4q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}