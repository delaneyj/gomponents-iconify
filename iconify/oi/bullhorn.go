package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullhorn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M6 0v6c.03.01.07 0 .09 0h.81c.05 0 .09-.04.09-.09V.1c0-.06-.04-.09-.09-.09h-.91zM5 .5L2.09 1.97c-.05.02-.13.03-.19.03H.09C.03 2 0 2.04 0 2.09V3.9c0 .06.04.09.09.09H1l1.03 2.72c.11.25.44.36.69.25c.25-.11.36-.44.25-.69l-.75-1.78c.03-.14.13-.22.28-.22v-.03L5 5.49v-5z"/>`),
		g.Group(children),
	)
}