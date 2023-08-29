package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselinePinOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 9l7 7h-3v4l-1 3l-1-3v-4H6v-3l3-3V9zm8-7v2l-2 1v5l3 3v2.461L12.27 9.73L9 6.46V5L7 4V2h10z"/><path fill="currentColor" d="M2.27 2.27L1 3.54L20.46 23l1.27-1.27L11 11z"/>`),
		g.Group(children),
	)
}