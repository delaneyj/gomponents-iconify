package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineRightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M186.83 125.17a4 4 0 0 1 0 5.66l-72 72a4 4 0 0 1-5.66-5.66L174.34 132H32a4 4 0 0 1 0-8h142.34l-65.17-65.17a4 4 0 0 1 5.66-5.66ZM216 36a4 4 0 0 0-4 4v176a4 4 0 0 0 8 0V40a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}