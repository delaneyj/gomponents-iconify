package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DvrSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14q.425 0 .713-.288T8 13q0-.425-.288-.713T7 12q-.425 0-.713.288T6 13q0 .425.288.713T7 14Zm0-4q.425 0 .713-.288T8 9q0-.425-.288-.713T7 8q-.425 0-.713.288T6 9q0 .425.288.713T7 10Zm2 4h9v-2H9v2Zm0-4h9V8H9v2ZM8 21v-2H2V3h20v16h-6v2H8Z"/>`),
		g.Group(children),
	)
}