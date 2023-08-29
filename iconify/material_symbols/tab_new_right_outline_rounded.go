package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabNewRightOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19V5v14v-.025V19Zm0 2q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v8.35q-.475-.175-.975-.263T19 13V5H5v14h8q0 .525.088 1.025t.262.975H5Zm14.175-1H16q-.425 0-.713-.288T15 19q0-.425.288-.713T16 18h3.175l-.9-.9Q18 16.825 18 16.412t.3-.712q.275-.275.7-.275t.7.275l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.275.275-.687.288T18.3 22.3q-.275-.275-.275-.7t.275-.7l.875-.9ZM11 13v3q0 .425.288.713T12 17q.425 0 .713-.288T13 16v-3h3q.425 0 .713-.288T17 12q0-.425-.288-.713T16 11h-3V8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8v3H8q-.425 0-.713.288T7 12q0 .425.288.713T8 13h3Z"/>`),
		g.Group(children),
	)
}