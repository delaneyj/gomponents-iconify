package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRearOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 4q0-.825.588-1.413T7 2h10q.825 0 1.413.588T19 4v12q0 .425-.288.713T18 17h-1V4H7v12.925L6.925 17H6q-.425 0-.713-.288T5 16V4Zm7 6q.825 0 1.413-.588T14 8q0-.825-.588-1.413T12 6q-.825 0-1.413.588T10 8q0 .825.588 1.413T12 10ZM9.55 20H6q-.425 0-.713-.288T5 19q0-.425.288-.713T6 18h3.55l-.4-.4q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l2.1 2.1q.3.3.3.7t-.3.7l-2.1 2.1q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l.4-.4ZM15 20q-.425 0-.713-.288T14 19q0-.425.288-.713T15 18h3q.425 0 .713.288T19 19q0 .425-.288.713T18 20h-3ZM12 8Z"/>`),
		g.Group(children),
	)
}