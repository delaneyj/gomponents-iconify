package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentMarginsTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2v-3.5a.5.5 0 0 1 1 0V18h6v-3.5a.5.5 0 0 1 1 0V18a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2v3.5a.5.5 0 0 1-1 0V2H7v3.5a.5.5 0 0 1-1 0V2Zm.5 11a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 1 0v5a.5.5 0 0 1-.5.5Zm7 0a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 1 0v5a.5.5 0 0 1-.5.5Z"/>`),
		g.Group(children),
	)
}