package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineUpThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M202.83 141.17a4 4 0 0 1-5.66 5.66L132 81.66V224a4 4 0 0 1-8 0V81.66l-65.17 65.17a4 4 0 0 1-5.66-5.66l72-72a4 4 0 0 1 5.66 0ZM216 36H40a4 4 0 0 0 0 8h176a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}