package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretCircleRightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm26.83-94.83a4 4 0 0 1 0 5.66l-40 40a4 4 0 0 1-5.66-5.66L146.34 128l-37.17-37.17a4 4 0 0 1 5.66-5.66Z"/>`),
		g.Group(children),
	)
}