package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M7.5.5v14m7-7.005H.5m13 0a6.006 6.006 0 0 1-6 6.005c-3.313 0-6-2.694-6-6.005a5.999 5.999 0 0 1 6-5.996a6 6 0 0 1 6 5.996Z"/>`),
		g.Group(children),
	)
}