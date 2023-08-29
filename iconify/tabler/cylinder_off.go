package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CylinderOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.23 5.233C5.08 5.478 5 5.735 5 6c0 1.131 1.461 2.117 3.62 2.628m4.357.343C16.381 8.767 19 7.515 19 6c0-1.657-3.134-3-7-3c-1.645 0-3.158.243-4.353.65"/><path d="M5 6v12c0 1.657 3.134 3 7 3c3.245 0 5.974-.946 6.767-2.23M19 15V6M3 3l18 18"/></g>`),
		g.Group(children),
	)
}