package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HangingProtocol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M26.37 20.52L17 15.41v-1.55A4 4 0 1 0 12 10h2a2 2 0 1 1 4 .34A2.08 2.08 0 0 1 15.89 12a.89.89 0 0 0-.89.89v2.52l-9.37 5.11A3.1 3.1 0 0 0 4.25 22a2.83 2.83 0 0 0 2.56 4h18.38a2.83 2.83 0 0 0 2.56-4a3.1 3.1 0 0 0-1.38-1.48zM25.19 24H6.81a.81.81 0 0 1-.81-.81a1 1 0 0 1 .52-.88L16 17.14l9.48 5.17a1 1 0 0 1 .52.88a.81.81 0 0 1-.81.81z" fill="currentColor"/>`),
		g.Group(children),
	)
}