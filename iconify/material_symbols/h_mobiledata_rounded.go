package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HMobiledataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 13v3q0 .425-.288.713T8 17q-.425 0-.713-.288T7 16V8q0-.425.288-.713T8 7q.425 0 .713.288T9 8v3h6V8q0-.425.288-.713T16 7q.425 0 .713.288T17 8v8q0 .425-.288.713T16 17q-.425 0-.713-.288T15 16v-3H9Z"/>`),
		g.Group(children),
	)
}