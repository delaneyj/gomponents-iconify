package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PipelineThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a2 2 0 1 1 4 0v16a2 2 0 1 1-4 0V8Zm24 0a2 2 0 1 1 4 0v16a2 2 0 1 1-4 0V8Zm-2 1H8v14h16V9Z"/>`),
		g.Group(children),
	)
}