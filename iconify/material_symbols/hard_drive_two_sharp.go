package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveTwoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 19q.425 0 .713-.288T15 18q0-.425-.288-.713T14 17q-.425 0-.713.288T13 18q0 .425.288.713T14 19Zm3 0q.425 0 .713-.288T18 18q0-.425-.288-.713T17 17q-.425 0-.713.288T16 18q0 .425.288.713T17 19ZM3 13V2h18v11H3Zm0 9v-7h18v7H3Z"/>`),
		g.Group(children),
	)
}