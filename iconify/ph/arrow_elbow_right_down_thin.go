package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightDownThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m226.83 162.83l-48 48a4 4 0 0 1-5.66 0l-48-48a4 4 0 1 1 5.66-5.66L172 198.34V68H32a4 4 0 0 1 0-8h144a4 4 0 0 1 4 4v134.34l41.17-41.17a4 4 0 0 1 5.66 5.66Z"/>`),
		g.Group(children),
	)
}