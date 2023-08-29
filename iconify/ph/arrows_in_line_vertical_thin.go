package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInLineVerticalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 128a4 4 0 0 1-4 4H40a4 4 0 0 1 0-8h176a4 4 0 0 1 4 4Zm-94.83-29.17a4 4 0 0 0 5.66 0l32-32a4 4 0 1 0-5.66-5.66L132 86.34V16a4 4 0 0 0-8 0v70.34L98.83 61.17a4 4 0 0 0-5.66 5.66Zm5.66 58.34a4 4 0 0 0-5.66 0l-32 32a4 4 0 0 0 5.66 5.66L124 169.66V240a4 4 0 0 0 8 0v-70.34l25.17 25.17a4 4 0 0 0 5.66-5.66Z"/>`),
		g.Group(children),
	)
}