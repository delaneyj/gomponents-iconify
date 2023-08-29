package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24.5 6a3.5 3.5 0 1 1 0 7a3.5 3.5 0 0 1 0-7Zm-5.41 4.5A5.502 5.502 0 0 0 30 9.5a5.5 5.5 0 0 0-10.91-1H19a4 4 0 0 0-4 4v7a2 2 0 0 1-2 2h-.09A5.502 5.502 0 0 0 2 22.5a5.5 5.5 0 0 0 10.91 1H13a4 4 0 0 0 4-4v-7a2 2 0 0 1 2-2h.09ZM11 22.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z"/>`),
		g.Group(children),
	)
}