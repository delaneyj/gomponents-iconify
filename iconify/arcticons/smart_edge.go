package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartEdge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.5 24a7.12 7.12 0 0 1-2.306 5.264H5.806C4.384 27.958 3.5 26.084 3.5 24s.884-3.958 2.306-5.264h36.388A7.12 7.12 0 0 1 44.5 24Z"/>`),
		g.Group(children),
	)
}