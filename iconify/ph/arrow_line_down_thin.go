package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineDownThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M53.17 114.83a4 4 0 0 1 5.66-5.66L124 174.34V32a4 4 0 0 1 8 0v142.34l65.17-65.17a4 4 0 1 1 5.66 5.66l-72 72a4 4 0 0 1-5.66 0ZM216 212H40a4 4 0 0 0 0 8h176a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}