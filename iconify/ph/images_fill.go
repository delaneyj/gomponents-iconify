package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 40H72a16 16 0 0 0-16 16v16H40a16 16 0 0 0-16 16v112a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16v-16h16a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm-44 32a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm12 128H40V88h16v80a16 16 0 0 0 16 16h112Zm32-32H72v-36l36-36l49.66 49.66a8 8 0 0 0 11.31 0L194.63 120L216 141.38V168Z"/>`),
		g.Group(children),
	)
}