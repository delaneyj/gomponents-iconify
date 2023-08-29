package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EMobiledataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17q-.425 0-.713-.288T8 16V8q0-.425.288-.713T9 7h6q.425 0 .713.288T16 8q0 .425-.288.713T15 9h-5v2h5q.425 0 .713.288T16 12q0 .425-.288.713T15 13h-5v2h5q.425 0 .713.288T16 16q0 .425-.288.713T15 17H9Z"/>`),
		g.Group(children),
	)
}