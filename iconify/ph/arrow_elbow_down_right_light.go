package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowDownRightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m212.24 180.24l-48 48a6 6 0 0 1-8.48-8.48L193.51 182H64a6 6 0 0 1-6-6V32a6 6 0 0 1 12 0v138h123.51l-37.75-37.76a6 6 0 1 1 8.48-8.48l48 48a6 6 0 0 1 0 8.48Z"/>`),
		g.Group(children),
	)
}