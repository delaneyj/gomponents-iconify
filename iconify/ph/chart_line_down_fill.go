package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 208a8 8 0 0 1-8 8H32a8 8 0 0 1-8-8V48a8 8 0 0 1 16 0v12.69l56 56l26.34-26.35a8 8 0 0 1 11.32 0L180 136.69l14.34-14.35A8 8 0 0 1 208 128v40a8 8 0 0 1-8 8h-40a8 8 0 0 1-5.66-13.66L168.69 148L128 107.31l-26.34 26.35a8 8 0 0 1-11.32 0L40 83.31V200h184a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}