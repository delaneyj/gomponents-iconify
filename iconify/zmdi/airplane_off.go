package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m235 152l170 107v42l-67-21l-167-167V35q0-14 9-23t22.5-9t23 9t9.5 23v117zM21 72l27-27l336 336l-27 27l-122-122v79l42 32v32l-74-21l-75 21v-32l43-32V248L0 301v-42l128-80z"/>`),
		g.Group(children),
	)
}