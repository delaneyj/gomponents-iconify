package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColorResetOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.25 16.4l-1.55-1.5q.15-.425.225-.862T18 13.1q0-1.175-.438-2.225T16.25 9L12 4.8L9.8 6.95l-1.4-1.4L12 2l5.65 5.55q1.1 1.05 1.725 2.488T20 13.1q0 .9-.2 1.725t-.55 1.575Zm.55 6.2l-3.1-3.1q-1.025.725-2.2 1.113T12 21q-3.325 0-5.663-2.313T4 13.1q0-1.275.4-2.45T5.6 8.4L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM12 19q.9 0 1.713-.25t1.537-.7L7 9.85q-.525.8-.763 1.6T6 13.1q0 2.45 1.75 4.175T12 19Zm-.925-5.1Zm2.75-2.9Z"/>`),
		g.Group(children),
	)
}