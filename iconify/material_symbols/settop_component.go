package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettopComponent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 17V7h20v10H2Zm3-4h6v-2H5v2Zm9 0q.425 0 .713-.288T15 12q0-.425-.288-.713T14 11q-.425 0-.713.288T13 12q0 .425.288.713T14 13Zm3 0q.425 0 .713-.288T18 12q0-.425-.288-.713T17 11q-.425 0-.713.288T16 12q0 .425.288.713T17 13Z"/>`),
		g.Group(children),
	)
}