package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.09 10.5H19a2 2 0 0 0-2 2v7a4 4 0 0 1-4 4h-.09A5.502 5.502 0 0 1 2 22.5a5.5 5.5 0 0 1 10.91-1H13a2 2 0 0 0 2-2v-7a4 4 0 0 1 4-4h.09A5.502 5.502 0 0 1 30 9.5a5.5 5.5 0 0 1-10.91 1Z"/>`),
		g.Group(children),
	)
}