package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkAsUnreadOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17q-.825 0-1.413-.588T2 15V7.15q0-.375.213-.738T2.8 5.85l6.8-3.4q.425-.2.9-.213t.9.213l6.65 3.4q.3.15.513.475T18.85 7h-2.925L10.5 4.25L4 7.475V17Zm3 4q-.825 0-1.413-.588T5 19v-9q0-.825.588-1.413T7 8h13q.825 0 1.413.588T22 10v9q0 .825-.588 1.413T20 21H7Zm5.575-6.125L7 12v7h13v-7l-5.575 2.875Q14 15.1 13.5 15.1t-.925-.225Zm.925-1.525L20 10H7l6.5 3.35ZM20 10H7h13Z"/>`),
		g.Group(children),
	)
}