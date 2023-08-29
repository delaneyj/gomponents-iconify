package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMissed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m354 85l30 30l-192 192L43 158v98H0V85h171v43H73l119 119z"/>`),
		g.Group(children),
	)
}