package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QueueLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M34 64a6 6 0 0 1 6-6h176a6 6 0 0 1 0 12H40a6 6 0 0 1-6-6Zm102 58H40a6 6 0 0 0 0 12h96a6 6 0 0 0 0-12Zm0 64H40a6 6 0 0 0 0 12h96a6 6 0 0 0 0-12Zm110-26a6 6 0 0 1-2.82 5.09l-64 40A6 6 0 0 1 170 200v-80a6 6 0 0 1 9.18-5.09l64 40A6 6 0 0 1 246 160Zm-17.32 0L182 130.83v58.34Z"/>`),
		g.Group(children),
	)
}