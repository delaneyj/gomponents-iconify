package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m299 320l49-49l-105-104l-85 85L0 94l30-30l128 128l85-85l135 134l49-49v128H299z"/>`),
		g.Group(children),
	)
}