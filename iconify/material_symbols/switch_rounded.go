package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 17q.425 0 .713-.288T16 16V8q0-.425-.288-.713T15 7H9q-.425 0-.713.288T8 8v8q0 .425.288.713T9 17h6Zm-5-2V9h4v6h-4Zm-5 6q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm7-9q.425 0 .713-.288T13 11q0-.425-.288-.713T12 10q-.425 0-.713.288T11 11q0 .425.288.713T12 12Z"/>`),
		g.Group(children),
	)
}