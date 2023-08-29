package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplayOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 17V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v12q0 .825-.588 1.413T20 19h-3v-2h3V5H4v12h3v2H4q-.825 0-1.413-.588T2 17Zm10-5Zm-4.725 7.725L11.3 15.7q.3-.3.7-.3t.7.3l4.025 4.025q.35.35.163.813T16.2 21H7.8q-.5 0-.687-.463t.162-.812Z"/>`),
		g.Group(children),
	)
}