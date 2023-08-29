package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvertColorsOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.275 16.45L12 9.175V4.8L9.775 6.975l-1.4-1.4L10.95 3.05q.225-.225.5-.337T12 2.6q.275 0 .55.113t.5.337l4.6 4.525q1.2 1.2 1.775 2.588T20 13.1q0 .95-.2 1.813t-.525 1.537ZM12 21q-3.325 0-5.663-2.288T4 13.1q0-1.275.4-2.45T5.6 8.4L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-2.4-2.4q-1.05.775-2.287 1.137T12 21Zm0-2v-4.175l-5-5q-.525.8-.763 1.613T6 13.1q0 2.5 1.75 4.2T12 19Z"/>`),
		g.Group(children),
	)
}