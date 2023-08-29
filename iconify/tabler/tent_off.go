package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 14l4 6h5m-2.863-6.868L12 4l-1.44 2.559M9.12 9.122L3 20h6l4-6M3 3l18 18"/>`),
		g.Group(children),
	)
}