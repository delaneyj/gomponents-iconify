package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm77.58 50.59L132 121.07v-85a92.07 92.07 0 0 1 73.58 42.52ZM124 36.09v89.6l-77.58 44.79A92 92 0 0 1 124 36.09ZM128 220a92 92 0 0 1-77.58-42.59l159.16-91.89A92 92 0 0 1 128 220Z"/>`),
		g.Group(children),
	)
}