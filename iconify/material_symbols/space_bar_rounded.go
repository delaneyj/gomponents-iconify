package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceBarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 15q-.825 0-1.413-.588T4 13v-3q0-.425.288-.713T5 9q.425 0 .713.288T6 10v3h12v-3q0-.425.288-.713T19 9q.425 0 .713.288T20 10v3q0 .825-.588 1.413T18 15H6Z"/>`),
		g.Group(children),
	)
}