package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VisualRecognition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="23" cy="12" r="2" fill="currentColor"/><path fill="currentColor" d="M28 5H16.24A8 8 0 1 0 6 16.92V27a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2ZM4 10a6 6 0 0 1 11.19-3H8v2h7.91a6.64 6.64 0 0 1 .09 1a6.64 6.64 0 0 1-.09 1H10v2h5.19A6 6 0 0 1 4 10Zm24 17H8l5-5l1.59 1.59a2 2 0 0 0 2.82 0L23 18l5 5Zm0-6.83l-3.59-3.59a2 2 0 0 0-2.82 0L16 22.17l-1.59-1.59a2 2 0 0 0-2.82 0L8 24.17v-6.43a8.24 8.24 0 0 0 2 .26a8 8 0 0 0 8-8a7.9 7.9 0 0 0-.59-3H28Z"/>`),
		g.Group(children),
	)
}