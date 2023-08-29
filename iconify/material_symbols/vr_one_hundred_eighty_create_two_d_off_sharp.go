package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrOneHundredEightyCreateTwoDOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17.75q-2.675-.675-4.338-2.838T2 10q0-1.125.3-2.15t.825-1.925l-2.5-2.5l1.4-1.4L21.2 21.2l-1.4 1.4l-.6-.6H10v-9.2l-2-1.975v6.925ZM10 2q2.75 0 4.913 1.663T17.75 8h-6.925L5.95 3.1q.925-.55 1.95-.825T10 2Zm2.825 8H22v9.175L12.825 10Zm3.375 9l-.7-.725l-1.425-1.4L12.5 19h3.7Z"/>`),
		g.Group(children),
	)
}