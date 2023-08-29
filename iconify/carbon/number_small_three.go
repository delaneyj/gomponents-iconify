package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSmallThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 11h-4v2h4v2h-3v2h3v2h-4v2h4a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}