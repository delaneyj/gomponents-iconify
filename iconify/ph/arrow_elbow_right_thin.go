package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 80v72a4 4 0 0 1-8 0V89.66L122.83 194.83a4 4 0 0 1-5.66 0l-96-96a4 4 0 0 1 5.66-5.66L120 186.34L222.34 84H160a4 4 0 0 1 0-8h72a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}