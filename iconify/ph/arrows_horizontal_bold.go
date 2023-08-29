package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m240.49 136.49l-32 32a12 12 0 0 1-17-17L203 140H53l11.52 11.51a12 12 0 0 1-17 17l-32-32a12 12 0 0 1 0-17l32-32a12 12 0 1 1 17 17L53 116h150l-11.52-11.51a12 12 0 0 1 17-17l32 32a12 12 0 0 1 .01 17Z"/>`),
		g.Group(children),
	)
}