package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInCardinalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M93.17 66.83a4 4 0 0 1 5.66-5.66L124 86.34V24a4 4 0 0 1 8 0v62.34l25.17-25.17a4 4 0 1 1 5.66 5.66l-32 32a4 4 0 0 1-5.66 0Zm37.66 90.34a4 4 0 0 0-5.66 0l-32 32a4 4 0 0 0 5.66 5.66L124 169.66V232a4 4 0 0 0 8 0v-62.34l25.17 25.17a4 4 0 0 0 5.66-5.66ZM232 124h-62.34l25.17-25.17a4 4 0 1 0-5.66-5.66l-32 32a4 4 0 0 0 0 5.66l32 32a4 4 0 0 0 5.66-5.66L169.66 132H232a4 4 0 0 0 0-8Zm-133.17 1.17l-32-32a4 4 0 0 0-5.66 5.66L86.34 124H24a4 4 0 0 0 0 8h62.34l-25.17 25.17a4 4 0 0 0 5.66 5.66l32-32a4 4 0 0 0 0-5.66Z"/>`),
		g.Group(children),
	)
}