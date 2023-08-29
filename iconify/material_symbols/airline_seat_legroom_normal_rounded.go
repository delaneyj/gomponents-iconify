package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomNormalRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 14H8q-.825 0-1.413-.588T6 12V3h6v6h5q.825 0 1.413.588T19 11v7h1.5q.65 0 1.075.425T22 19.5q0 .65-.425 1.075T20.5 21H17q-.425 0-.713-.288T16 20v-6Zm-3 3H5q-.825 0-1.413-.588T3 15V4q0-.425.288-.713T4 3q.425 0 .713.288T5 4v11h8q.425 0 .713.288T14 16q0 .425-.288.713T13 17Z"/>`),
		g.Group(children),
	)
}