package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18m-1.5-8.428L18 14m-2 2l-4 4l-7.5-7.428a5 5 0 0 1-1.288-5.068A4.976 4.976 0 0 1 5 5m3-1c1.56 0 3.05.727 4 2a5 5 0 1 1 7.5 6.572"/>`),
		g.Group(children),
	)
}