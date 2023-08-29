package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 56v144a12 12 0 0 1-24 0v-60H68v60a12 12 0 0 1-24 0V56a12 12 0 0 1 24 0v60h120V56a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}