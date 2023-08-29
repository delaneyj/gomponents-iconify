package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GirlRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 7.5q-.725 0-1.238-.513T10.25 5.75q0-.725.513-1.238T12 4q.725 0 1.238.513t.512 1.237q0 .725-.513 1.238T12 7.5ZM11 20q-.425 0-.713-.288T10 19v-3h-.55q-.525 0-.838-.425T8.5 14.65l1.875-5.025q.2-.5.637-.813T12 8.5q.55 0 .988.313t.637.812L15.5 14.65q.2.5-.113.925T14.55 16H14v3q0 .425-.288.713T13 20h-2Z"/>`),
		g.Group(children),
	)
}