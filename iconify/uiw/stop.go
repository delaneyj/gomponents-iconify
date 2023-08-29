package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0C4.478 0 0 4.478 0 10s4.478 10 10 10s10-4.478 10-10S15.522 0 10 0Zm5.241 16.44L3.561 4.759c.356-.44.758-.842 1.198-1.199l11.68 11.681a8.335 8.335 0 0 1-1.198 1.199Z"/>`),
		g.Group(children),
	)
}