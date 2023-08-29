package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M108 48v160a4 4 0 0 1-8 0v-76H33.66l25.17 25.17a4 4 0 0 1-5.66 5.66l-32-32a4 4 0 0 1 0-5.66l32-32a4 4 0 0 1 5.66 5.66L33.66 124H100V48a4 4 0 0 1 8 0Zm126.83 77.17l-32-32a4 4 0 0 0-5.66 5.66L222.34 124H156V48a4 4 0 0 0-8 0v160a4 4 0 0 0 8 0v-76h66.34l-25.17 25.17a4 4 0 0 0 5.66 5.66l32-32a4 4 0 0 0 0-5.66Z"/>`),
		g.Group(children),
	)
}