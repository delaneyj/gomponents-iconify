package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm60.59 110.55l-40 28A8 8 0 0 1 136 156v-28a8 8 0 0 1-3.41 6.55l-40 28A8 8 0 0 1 80 156v-56a8 8 0 0 1 12.59-6.55l40 28A8 8 0 0 1 136 128v-28a8 8 0 0 1 12.59-6.55l40 28a8 8 0 0 1 0 13.1Z"/>`),
		g.Group(children),
	)
}