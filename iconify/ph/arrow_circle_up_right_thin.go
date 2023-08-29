package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpRightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm36-124v48a4 4 0 0 1-8 0v-38.34l-57.17 57.17a4 4 0 0 1-5.66-5.66L150.34 100H112a4 4 0 0 1 0-8h48a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}