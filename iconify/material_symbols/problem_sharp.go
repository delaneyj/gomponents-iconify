package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProblemSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 16q.425 0 .713-.288T8 15q0-.425-.288-.713T7 14q-.425 0-.713.288T6 15q0 .425.288.713T7 16Zm-1-3h2V8H6v5Zm4 2h8v-2h-8v2Zm0-4h8V9h-8v2Zm-8 9V4h20v16H2Z"/>`),
		g.Group(children),
	)
}