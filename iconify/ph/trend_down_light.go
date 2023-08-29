package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendDownLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M238 136v64a6 6 0 0 1-6 6h-64a6 6 0 0 1 0-12h49.52L136 112.49l-35.76 35.75a6 6 0 0 1-8.48 0l-72-72a6 6 0 0 1 8.48-8.48L96 135.51l35.76-35.75a6 6 0 0 1 8.48 0L226 185.52V136a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}