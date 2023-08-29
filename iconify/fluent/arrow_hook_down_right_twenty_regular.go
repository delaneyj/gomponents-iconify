package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowHookDownRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 9a5 5 0 0 1 5-5h4.5a.5.5 0 0 1 0 1H9a4 4 0 1 0 0 8h5.293l-2.7-2.7a.5.5 0 1 1 .708-.706l3.539 3.539a.5.5 0 0 1 .125.497a.499.499 0 0 1-.135.247l-3.533 3.533a.5.5 0 0 1-.707-.707L14.293 14H9a5 5 0 0 1-5-5Z"/>`),
		g.Group(children),
	)
}