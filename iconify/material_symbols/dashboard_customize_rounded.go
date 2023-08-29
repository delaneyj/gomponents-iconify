package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardCustomizeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 11q-.425 0-.713-.288T3 10V4q0-.425.288-.713T4 3h6q.425 0 .713.288T11 4v6q0 .425-.288.713T10 11H4Zm10 0q-.425 0-.713-.288T13 10V4q0-.425.288-.713T14 3h6q.425 0 .713.288T21 4v6q0 .425-.288.713T20 11h-6ZM4 21q-.425 0-.713-.288T3 20v-6q0-.425.288-.713T4 13h6q.425 0 .713.288T11 14v6q0 .425-.288.713T10 21H4Zm13 0q-.425 0-.713-.288T16 20v-2h-2.025q-.425 0-.7-.288T13 17q0-.425.288-.713T14 16h2v-2.025q0-.425.288-.7T17 13q.425 0 .713.288T18 14v2h2.025q.425 0 .7.288T21 17q0 .425-.288.713T20 18h-2v2.025q0 .425-.288.7T17 21Z"/>`),
		g.Group(children),
	)
}