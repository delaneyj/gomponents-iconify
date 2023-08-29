package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesExt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-1.65 0-2.825-1.175T2 17q0-1.65 1.175-2.825T6 13q1.65 0 2.825 1.175T10 17q0 1.65-1.175 2.825T6 21Zm12 0q-1.65 0-2.825-1.175T14 17q0-1.65 1.175-2.825T18 13q1.65 0 2.825 1.175T22 17q0 1.65-1.175 2.825T18 21Zm-6-6.1q-.8 0-1.35-.55T10.1 13q0-.8.55-1.35T12 11.1q.8 0 1.35.55T13.9 13q0 .8-.55 1.35T12 14.9Zm0-4.9q-1.65 0-2.825-1.175T8 6q0-1.65 1.175-2.825T12 2q1.65 0 2.825 1.175T16 6q0 1.65-1.175 2.825T12 10Z"/>`),
		g.Group(children),
	)
}