package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomInRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 10.5h-1q-.425 0-.713-.288T6.5 9.5q0-.425.288-.713T7.5 8.5h1v-1q0-.425.288-.713T9.5 6.5q.425 0 .713.288t.287.712v1h1q.425 0 .713.288t.287.712q0 .425-.288.713t-.712.287h-1v1q0 .425-.288.713T9.5 12.5q-.425 0-.713-.288T8.5 11.5v-1Zm1 5.5q-2.725 0-4.612-1.888T3 9.5q0-2.725 1.888-4.612T9.5 3q2.725 0 4.612 1.888T16 9.5q0 1.1-.35 2.075T14.7 13.3l5.6 5.6q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-5.6-5.6q-.75.6-1.725.95T9.5 16Zm0-2q1.875 0 3.188-1.313T14 9.5q0-1.875-1.313-3.188T9.5 5Q7.625 5 6.312 6.313T5 9.5q0 1.875 1.313 3.188T9.5 14Z"/>`),
		g.Group(children),
	)
}