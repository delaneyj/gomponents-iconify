package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryHorizZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17q-.425 0-.713-.288T4 16v-2H2v-4h2V8q0-.425.288-.713T5 7h16q.425 0 .713.288T22 8v8q0 .425-.288.713T21 17H5Zm1-2h14V9H6v6Z"/>`),
		g.Group(children),
	)
}