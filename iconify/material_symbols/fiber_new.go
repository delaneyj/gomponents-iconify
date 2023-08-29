package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberNew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20q-.825 0-1.413-.588T1 18V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v12q0 .825-.588 1.413T21 20H3Zm.5-5h1.25v-3.5L7.3 15h1.2V9H7.25v3.5L4.75 9H3.5v6Zm6 0h4v-1.25H11v-1.1h2.5V11.4H11v-1.15h2.5V9h-4v6Zm6 0h4q.425 0 .713-.288T20.5 14V9h-1.25v4.5h-1.1V10H16.9v3.5h-1.15V9H14.5v5q0 .425.288.713T15.5 15Z"/>`),
		g.Group(children),
	)
}