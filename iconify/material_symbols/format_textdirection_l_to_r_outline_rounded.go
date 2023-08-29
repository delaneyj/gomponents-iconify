package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextdirectionLToROutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4v10q0 .425-.288.713T10 15q-.425 0-.713-.288T9 14v-4q-1.65 0-2.825-1.175T5 6q0-1.65 1.175-2.825T9 2h7q.425 0 .713.287T17 3q0 .425-.288.713T16 4h-1v10q0 .425-.288.713T14 15q-.425 0-.713-.288T13 14V4h-2ZM9 8V4q-.825 0-1.413.588T7 6q0 .825.588 1.413T9 8Zm0-2Zm8.2 13H4q-.425 0-.713-.288T3 18q0-.425.288-.713T4 17h13.2l-.9-.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l.9-.9Z"/>`),
		g.Group(children),
	)
}