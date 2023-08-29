package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendDownThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 136v64a4 4 0 0 1-4 4h-64a4 4 0 0 1 0-8h54.34L136 109.66l-37.17 37.17a4 4 0 0 1-5.66 0l-72-72a4 4 0 0 1 5.66-5.66L96 138.34l37.17-37.17a4 4 0 0 1 5.66 0L228 190.34V136a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}