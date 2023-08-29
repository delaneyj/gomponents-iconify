package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 56v144a4 4 0 0 1-8 0v-68H60v68a4 4 0 0 1-8 0V56a4 4 0 0 1 8 0v68h136V56a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}