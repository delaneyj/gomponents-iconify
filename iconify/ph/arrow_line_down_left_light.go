package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineDownLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 40a6 6 0 0 1-6 6H40a6 6 0 0 1 0-12h176a6 6 0 0 1 6 6Zm-50.24 43.76L70 185.52V104a6 6 0 0 0-12 0v96a6 6 0 0 0 6 6h96a6 6 0 0 0 0-12H78.48L180.24 92.24a6 6 0 0 0-8.48-8.48Z"/>`),
		g.Group(children),
	)
}