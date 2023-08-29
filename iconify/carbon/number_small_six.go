package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSmallSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 21h-2a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h3v2h-3v2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2Zm-2-4v2h2v-2Z"/>`),
		g.Group(children),
	)
}