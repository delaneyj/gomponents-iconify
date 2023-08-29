package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M116 48v160a12 12 0 0 1-24 0v-68H53l11.52 11.51a12 12 0 0 1-17 17l-32-32a12 12 0 0 1 0-17l32-32a12 12 0 1 1 17 17L53 116h39V48a12 12 0 0 1 24 0Zm124.49 71.51l-32-32a12 12 0 0 0-17 17L203 116h-39V48a12 12 0 0 0-24 0v160a12 12 0 0 0 24 0v-68h39l-11.52 11.51a12 12 0 0 0 17 17l32-32a12 12 0 0 0 .01-17Z"/>`),
		g.Group(children),
	)
}