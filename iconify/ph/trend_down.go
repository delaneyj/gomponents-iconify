package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 136v64a8 8 0 0 1-8 8h-64a8 8 0 0 1 0-16h44.69L136 115.31l-34.34 34.35a8 8 0 0 1-11.32 0l-72-72a8 8 0 0 1 11.32-11.32L96 132.69l34.34-34.35a8 8 0 0 1 11.32 0L224 180.69V136a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}