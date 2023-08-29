package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.13 104.13 0 0 0 128 24Zm36.44 110.66l-48 32A8.05 8.05 0 0 1 112 168a8 8 0 0 1-8-8V96a8 8 0 0 1 12.44-6.66l48 32a8 8 0 0 1 0 13.32Z"/>`),
		g.Group(children),
	)
}