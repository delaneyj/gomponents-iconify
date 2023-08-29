package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M206 56v144a6 6 0 0 1-12 0v-66H62v66a6 6 0 0 1-12 0V56a6 6 0 0 1 12 0v66h132V56a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}