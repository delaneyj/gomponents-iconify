package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReorderRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19q-.425 0-.713-.288T3 18q0-.425.288-.713T4 17h16q.425 0 .713.288T21 18q0 .425-.288.713T20 19H4Zm0-4q-.425 0-.713-.288T3 14q0-.425.288-.713T4 13h16q.425 0 .713.288T21 14q0 .425-.288.713T20 15H4Zm0-4q-.425 0-.713-.288T3 10q0-.425.288-.713T4 9h16q.425 0 .713.288T21 10q0 .425-.288.713T20 11H4Zm0-4q-.425 0-.713-.288T3 6q0-.425.288-.713T4 5h16q.425 0 .713.288T21 6q0 .425-.288.713T20 7H4Z"/>`),
		g.Group(children),
	)
}