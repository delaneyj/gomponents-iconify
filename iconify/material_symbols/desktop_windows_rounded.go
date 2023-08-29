package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopWindowsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 19v-2H4q-.825 0-1.413-.588T2 15V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v10q0 .825-.588 1.413T20 17h-6v2h1q.425 0 .713.288T16 20q0 .425-.288.713T15 21H9q-.425 0-.713-.288T8 20q0-.425.288-.713T9 19h1Z"/>`),
		g.Group(children),
	)
}