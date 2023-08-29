package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PipelineThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 9a3 3 0 0 1 6 0h16a3 3 0 1 1 6 0v14a3 3 0 1 1-6 0H8a3 3 0 1 1-6 0V9Zm4 0a1 1 0 0 0-2 0v14a1 1 0 1 0 2 0V9Zm2 12h16V11H8v10ZM28 9a1 1 0 1 0-2 0v14a1 1 0 1 0 2 0V9Z"/>`),
		g.Group(children),
	)
}