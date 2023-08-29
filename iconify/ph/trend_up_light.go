package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendUpLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M238 56v64a6 6 0 0 1-12 0V70.48l-85.76 85.76a6 6 0 0 1-8.48 0L96 120.49l-67.76 67.75a6 6 0 0 1-8.48-8.48l72-72a6 6 0 0 1 8.48 0L136 143.51L217.52 62H168a6 6 0 0 1 0-12h64a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}