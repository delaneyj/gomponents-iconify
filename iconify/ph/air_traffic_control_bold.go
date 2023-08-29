package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirTrafficControlBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232.38 60.52A20 20 0 0 0 216 52h-76V28h12a12 12 0 0 0 0-24h-48a12 12 0 0 0 0 24h12v24H40a20 20 0 0 0-18.8 26.83l26.19 72A20.06 20.06 0 0 0 66.18 164H92v60a12 12 0 0 0 24 0v-60h24v60a12 12 0 0 0 24 0v-60h25.82a20.06 20.06 0 0 0 18.79-13.17l26.19-72a20 20 0 0 0-2.42-18.31ZM114 140l-11.62-64h51.24L142 140ZM45.71 76H78l11.64 64H69ZM187 140h-20.62L178 76h32.27Z"/>`),
		g.Group(children),
	)
}