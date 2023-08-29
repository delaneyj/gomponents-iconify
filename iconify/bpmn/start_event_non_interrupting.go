package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StartEventNonInterrupting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-dasharray="418.31 361.233" stroke-linecap="round" stroke-width="100" d="M1899 1023.995c21.545 567.431-598.38 1023.5-1133.58 835.92C217.327 1712.701-36.147 985.964 301.83 529.931C604.45 49.634 1373.576 22.397 1709.411 480.086c122.14 153.124 190.066 348.066 189.589 543.909z"/>`),
		g.Group(children),
	)
}