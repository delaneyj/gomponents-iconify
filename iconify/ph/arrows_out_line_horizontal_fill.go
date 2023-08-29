package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutLineHorizontalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M104 128a8 8 0 0 1-8 8H56v24a8 8 0 0 1-13.66 5.66l-32-32a8 8 0 0 1 0-11.32l32-32A8 8 0 0 1 56 96v24h40a8 8 0 0 1 8 8Zm141.66-5.66l-32-32A8 8 0 0 0 200 96v24h-40a8 8 0 0 0 0 16h40v24a8 8 0 0 0 13.66 5.66l32-32a8 8 0 0 0 0-11.32ZM128 32a8 8 0 0 0-8 8v176a8 8 0 0 0 16 0V40a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}