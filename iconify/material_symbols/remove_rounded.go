package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 13q-.425 0-.713-.288T5 12q0-.425.288-.713T6 11h12q.425 0 .713.288T19 12q0 .425-.288.713T18 13H6Z"/>`),
		g.Group(children),
	)
}