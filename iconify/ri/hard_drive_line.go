package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 10.938A8.004 8.004 0 0 0 11.938 4H5v6.938Zm0 2.013V20h14V4h-5.05A10.003 10.003 0 0 1 5 12.95ZM4 2h16a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1Zm11 14h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}