package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectStartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 12v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2ZM5 19V5H4q-.425 0-.713-.288T3 4q0-.425.288-.713T4 3h4q.425 0 .713.288T9 4q0 .425-.288.713T8 5H7v14h1q.425 0 .713.288T9 20q0 .425-.288.713T8 21H4q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h1ZM19 5V3q.825 0 1.413.588T21 5h-2Zm0 16v-2h2q0 .825-.588 1.413T19 21Z"/>`),
		g.Group(children),
	)
}