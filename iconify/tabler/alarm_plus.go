package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13a7 7 0 1 0 14 0a7 7 0 1 0-14 0m2-9L4.25 6M17 4l2.75 2M10 13h4m-2-2v4"/>`),
		g.Group(children),
	)
}