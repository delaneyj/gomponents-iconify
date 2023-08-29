package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15.5 18.5l-3 1.5l.5-3.5l-2-2l3-.5l1.5-3l1.5 3l3 .5l-2 2l.5 3.5zM4 4l7 7M9 4l3.5 3.5M4 9l3.5 3.5"/>`),
		g.Group(children),
	)
}