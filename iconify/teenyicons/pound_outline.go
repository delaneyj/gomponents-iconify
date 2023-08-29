package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoundOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5 14.5l1.103-.367A2.775 2.775 0 0 0 3.5 11.5V4.442A3.942 3.942 0 0 1 7.442.5h.865C9.934.5 11.396 1.49 12 3M3 13h1.084a6 6 0 0 1 2.683.633l.05.025a6 6 0 0 0 5.366 0L13.5 13M1 7.5h8"/>`),
		g.Group(children),
	)
}