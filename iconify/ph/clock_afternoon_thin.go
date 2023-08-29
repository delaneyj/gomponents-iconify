package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockAfternoonThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm60-92a4 4 0 0 1-4 4h-46.34l33.17 33.17a4 4 0 0 1-5.66 5.66l-40-40A4 4 0 0 1 128 124h56a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}