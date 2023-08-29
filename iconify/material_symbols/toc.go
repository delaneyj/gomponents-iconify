package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17v-2h14v2H3Zm0-4v-2h14v2H3Zm0-4V7h14v2H3Zm17 8q-.425 0-.713-.288T19 16q0-.425.288-.713T20 15q.425 0 .713.288T21 16q0 .425-.288.713T20 17Zm0-4q-.425 0-.713-.288T19 12q0-.425.288-.713T20 11q.425 0 .713.288T21 12q0 .425-.288.713T20 13Zm0-4q-.425 0-.713-.288T19 8q0-.425.288-.713T20 7q.425 0 .713.288T21 8q0 .425-.288.713T20 9Z"/>`),
		g.Group(children),
	)
}