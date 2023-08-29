package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LooksFourOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 13v3q0 .425.288.713T14 17q.425 0 .713-.288T15 16V8q0-.425-.288-.713T14 7q-.425 0-.713.288T13 8v3h-2V8q0-.425-.288-.713T10 7q-.425 0-.713.288T9 8v4q0 .425.288.713T10 13h3Zm-8 8q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}