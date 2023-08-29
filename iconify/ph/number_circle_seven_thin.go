package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleSevenThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm27.28-134.29a4 4 0 0 1 .48 3.66l-32 88A4 4 0 0 1 120 180a4.12 4.12 0 0 1-1.37-.24a4 4 0 0 1-2.39-5.13L146.29 92H104a4 4 0 0 1 0-8h48a4 4 0 0 1 3.28 1.71Z"/>`),
		g.Group(children),
	)
}