package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSmallSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.44 21h-2L17 13h-2v1h-2v-3h6v2l-2.56 8z"/>`),
		g.Group(children),
	)
}