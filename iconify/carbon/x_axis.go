package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XAxis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m23 20l-1.414 1.414L24.172 24H6V4H4v20a2.002 2.002 0 0 0 2 2h18.172l-2.586 2.586L23 30l5-5Z"/>`),
		g.Group(children),
	)
}