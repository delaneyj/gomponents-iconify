package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateCounterclockwise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 28V16a2.002 2.002 0 0 1 2-2h12a2.002 2.002 0 0 1 2 2v12a2.002 2.002 0 0 1-2 2H4a2.002 2.002 0 0 1-2-2zm2-12l-.001 12H16V16zM17 2l1.41 1.41L15.83 6H21a7.008 7.008 0 0 1 7 7v5h-2v-5a5.006 5.006 0 0 0-5-5h-5.17l2.58 2.59L17 12l-5-5z"/>`),
		g.Group(children),
	)
}