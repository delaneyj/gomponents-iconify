package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaceItemRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V9q0-.825.588-1.413T5 7h3q.425 0 .713.288T9 8q0 .425-.288.713T8 9H5v10h14V9h-3q-.425 0-.713-.288T15 8q0-.425.288-.713T16 7h3q.825 0 1.413.588T21 9v10q0 .825-.588 1.413T19 21H5Zm8-8.825l.9-.9q.275-.275.688-.275t.712.3q.275.275.275.7t-.275.7l-2.6 2.6q-.3.3-.7.3t-.7-.3l-2.6-2.6q-.275-.275-.287-.687T8.7 11.3q.275-.275.7-.275t.7.275l.9.875V1q0-.425.288-.713T12 0q.425 0 .713.288T13 1v11.175Z"/>`),
		g.Group(children),
	)
}