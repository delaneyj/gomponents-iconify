package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GMobiledataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17q-.825 0-1.413-.588T7 15V9q0-.825.588-1.413T9 7h6q.425 0 .713.288T16 8q0 .425-.288.713T15 9H9v6h5v-2h-1q-.425 0-.713-.288T12 12q0-.425.288-.713T13 11h2q.425 0 .713.288T16 12v3q0 .825-.588 1.413T14 17H9Z"/>`),
		g.Group(children),
	)
}