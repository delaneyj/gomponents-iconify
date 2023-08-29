package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackStarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 17.125l1.775 1.075q.15.1.288 0t.087-.275l-.475-2.025l1.55-1.325q.125-.125.075-.275t-.225-.175l-2.025-.175l-.825-1.925q-.05-.15-.225-.15t-.225.15l-.825 1.925l-2.025.175q-.175.025-.225.175t.075.275l1.55 1.325l-.475 2.025q-.05.175.088.275t.287 0L15 17.125ZM4 16q-.825 0-1.413-.588T2 14V4q0-.825.588-1.413T4 2h10q.825 0 1.413.588T16 4v1q0 .425-.288.713T15 6q-.425 0-.713-.288T14 5V4H4v10h1q.425 0 .713.288T6 15q0 .425-.288.713T5 16H4Zm6 6q-.825 0-1.413-.588T8 20V10q0-.825.588-1.413T10 8h10q.825 0 1.413.588T22 10v10q0 .825-.588 1.413T20 22H10Z"/>`),
		g.Group(children),
	)
}