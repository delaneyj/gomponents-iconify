package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M32 22v2H0V0h2v22zM30 2.5v6.802a.498.498 0 0 1-.859.342l-1.89-1.89l-9.89 9.887a.49.49 0 0 1-.72 0L13.001 14l-6.5 6.5l-3-3l9.14-9.141a.49.49 0 0 1 .72 0L17.001 12l7.25-7.25l-1.89-1.89a.497.497 0 0 1 .342-.86h6.822a.48.48 0 0 1 .48.48v.021V2.5z"/>`),
		g.Group(children),
	)
}