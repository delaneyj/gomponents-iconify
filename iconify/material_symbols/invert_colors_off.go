package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvertColorsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.275 16.45L12 9.175V4.8L9.775 6.975l-1.4-1.4L12 2l5.65 5.575q1.2 1.2 1.775 2.588T20 13.1q0 .95-.2 1.813t-.525 1.537Zm.525 6.15l-3.1-3.1q-1.05.775-2.288 1.137T12 21q-3.325 0-5.663-2.288T4 13.1q0-1.275.4-2.45T5.6 8.4L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM12 19v-4.175l-5-5q-.525.8-.763 1.613T6 13.1q0 2.5 1.75 4.2T12 19Z"/>`),
		g.Group(children),
	)
}