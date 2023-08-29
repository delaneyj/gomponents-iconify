package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBackCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm33.94-135.5a4 4 0 0 0-4.06.11L100 120.78V88a4 4 0 0 0-8 0v80a4 4 0 0 0 8 0v-32.78l57.88 36.17a4 4 0 0 0 2.12.61a4.06 4.06 0 0 0 1.94-.5A4 4 0 0 0 164 168V88a4 4 0 0 0-2.06-3.5ZM156 160.78L103.55 128L156 95.22Z"/>`),
		g.Group(children),
	)
}