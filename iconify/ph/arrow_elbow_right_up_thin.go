package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightUpThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M226.83 98.83a4 4 0 0 1-5.66 0L180 57.66V192a4 4 0 0 1-4 4H32a4 4 0 0 1 0-8h140V57.66l-41.17 41.17a4 4 0 1 1-5.66-5.66l48-48a4 4 0 0 1 5.66 0l48 48a4 4 0 0 1 0 5.66Z"/>`),
		g.Group(children),
	)
}