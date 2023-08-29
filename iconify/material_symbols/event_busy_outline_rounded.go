package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventBusyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.4L10.4 18q-.275.275-.7.275T9 18q-.275-.275-.275-.7T9 16.6l1.6-1.6L9 13.4q-.275-.275-.275-.7T9 12q.275-.275.7-.275t.7.275l1.6 1.6l1.6-1.6q.275-.275.7-.275T15 12q.275.275.275.7t-.275.7L13.4 15l1.6 1.6q.275.275.275.7T15 18q-.275.275-.7.275T13.6 18L12 16.4ZM5 22q-.825 0-1.413-.588T3 20V6q0-.825.588-1.413T5 4h1V3q0-.425.288-.713T7 2q.425 0 .713.288T8 3v1h8V3q0-.425.288-.713T17 2q.425 0 .713.288T18 3v1h1q.825 0 1.413.588T21 6v14q0 .825-.588 1.413T19 22H5Zm0-2h14V10H5v10ZM5 8h14V6H5v2Zm0 0V6v2Z"/>`),
		g.Group(children),
	)
}