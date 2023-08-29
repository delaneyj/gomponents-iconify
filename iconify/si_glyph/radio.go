package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 0h.938v4.93H3zm9.018 6.01v10h2.962v-10h-2.962zm2.003 4.011h-1.042V9h1.042v1.021zm0-2h-1.042V7h1.042v1.021zm-13.01 7.995h10v-10h-10v10zm4.497-.954a3.548 3.548 0 0 1-2.562-1.094h5.125a3.548 3.548 0 0 1-2.563 1.094zm3.521-3.093a3.54 3.54 0 0 1-.307 1.047h-6.43a3.57 3.57 0 0 1-.307-1.047h7.044zm-7.045-.953c.049-.37.147-.722.299-1.047H8.73c.153.325.252.677.301 1.047H1.984zm3.524-3.11c1.016 0 1.926.429 2.576 1.109H2.932a3.551 3.551 0 0 1 2.576-1.109z"/>`),
		g.Group(children),
	)
}