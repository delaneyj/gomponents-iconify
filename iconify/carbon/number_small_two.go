package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSmallTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 21h-6v-4a2 2 0 0 1 2-2h2v-2h-4v-2h4a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2v2h4Z"/>`),
		g.Group(children),
	)
}