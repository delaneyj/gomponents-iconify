package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1435 718q0 111-75.5 192.5T1174 1000v1H304v-1q-7 1-21 1q-117 0-200-83T0 718q0-74 37-139t101-103q-8-32-8-62q0-117 83-200t200-83q103 0 186 70q43-91 128.5-145.5T913 1q144 0 246 102t102 246q0 55-16 103q85 29 137.5 103t52.5 163z"/>`),
		g.Group(children),
	)
}