package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MasterOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20h18v2H3zm18-1H3L2.147 7.81A2 2 0 1 1 5 6a1.914 1.914 0 0 1-.024.3l2.737 2.189l2.562-4.486A1.948 1.948 0 0 1 10 3a2 2 0 0 1 4 0a1.946 1.946 0 0 1-.276 1.004l2.558 4.485l2.741-2.19A1.906 1.906 0 0 1 19 6a2 2 0 1 1 2.853 1.81ZM4.92 17h14.16l.73-8.77l-4.106 3.281L12 5.017l-3.71 6.494l-4.1-3.28Z"/>`),
		g.Group(children),
	)
}