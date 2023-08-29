package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneBluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m250 187l-15-15l59-60l-59-60l15-15l49 49V5h10l61 61l-46 46l46 46l-61 61h-10v-81zm70-141v40l20-20zm0 92v40l20-20zm43 177q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 432q-99 0-182.5-48.5t-132-132T0 69q0-8 6.5-14.5T21 48h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22l-47 47q47 93 141 141l47-47q9-9 22-5q36 12 76 12z"/>`),
		g.Group(children),
	)
}